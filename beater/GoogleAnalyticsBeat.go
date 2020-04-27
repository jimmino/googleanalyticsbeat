package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"

	"github.com/jimmino/googleanalyticsbeat/config"
	"github.com/jimmino/googleanalyticsbeat/ga"
)

var debugf = logp.MakeDebug("gabeat")

type gaDataRetriever func(gaConfig config.GoogleAnalyticsConfig) (gaDataPoints []ga.GABeatDataPoint, err error)

// googleanalyticsbeat configuration.
type googleanalyticsbeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New creates an instance of googleanalyticsbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &googleanalyticsbeat{
		done:   make(chan struct{}),
		config: c,
	}
	logp.Info("Config: %s", bt.config)
	return bt, nil
}

// Run starts googleanalyticsbeat.
func (bt *googleanalyticsbeat) Run(b *beat.Beat) error {
	return runFunctionally(bt, b, ga.GetGAReportData)
}

func runFunctionally(bt *googleanalyticsbeat, b *beat.Beat, dataFunc gaDataRetriever) error {
	logp.Info("googleanalyticsbeat is running! Hit CTRL-C to stop it.")
	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(bt.config.Period)
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		} //end select
		beatOnce(bt.client, b.Info.Name, bt.config.Googleanalytics, dataFunc)
	} //end for
} //end func

func beatOnce(client beat.Client, beatName string, gaConfig config.GoogleAnalyticsConfig, dataFunc gaDataRetriever) {
	GAData, err := dataFunc(gaConfig)
	if err == nil {
		publishToElastic(client, beatName, GAData)
	} else {
		logp.Err("gadata was null, not publishing: %v", err)
	}

}

func makeEvent(beatType string, GAData []ga.GABeatDataPoint) beat.Event {
	//event := common.MapStr{
	event := beat.Event{
		//"@timestamp": common.Time(time.Now()),
		Timestamp: time.Now(),
		Fields: common.MapStr{
			"type":  beatType,
			"count": 1, //The number of transactions that this event represents
		},
	}
	for _, gaDataPoint := range GAData {
		gaDataName := gaDataPoint.DimensionName + "_" + gaDataPoint.MetricName
		event.PutValue(gaDataName, gaDataPoint.Value)
	}
	return event
}

func publishToElastic(client beat.Client, beatType string, GAData []ga.GABeatDataPoint) {
	event := makeEvent(beatType, GAData)
	//succeeded := client.PublishEvent(event)
	client.Publish(event)
	logp.Info("Event sent")
	/*if !succeeded {
		logp.Err("Publisher couldn't publish event to Elastic")
	}*/
}

// Stop stops googleanalyticsbeat.
func (bt *googleanalyticsbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
