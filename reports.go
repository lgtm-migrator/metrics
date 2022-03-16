package metrics

func SlowSubscriberEventsDropped(amount int, tags ...Tags) {
	clientMux.RLock()
	defer clientMux.RUnlock()

	if client == nil {
		return
	}

	tagSpec := joinTags(tags...)

	client.Count("pubsub.slow_subscriber.events_dropped", int64(amount), tagSpec, 1)
}

func SpotTradesBatchSubmitted(size int, tags ...Tags) {
	clientMux.RLock()
	defer clientMux.RUnlock()

	if client == nil {
		return
	}

	tagSpec := joinTags(tags...)

	client.Count("events.spot_trades_batch.size", int64(size), tagSpec, 1)
}

func DerivativeTradesBatchSubmitted(size int, tags ...Tags) {
	clientMux.RLock()
	defer clientMux.RUnlock()

	if client == nil {
		return
	}

	tagSpec := joinTags(tags...)

	client.Count("events.derivative_trades_batch.size", int64(size), tagSpec, 1)
}

func IndexPriceUpdatesBatchSubmitted(size int, tags ...Tags) {
	clientMux.RLock()
	defer clientMux.RUnlock()

	if client == nil {
		return
	}

	tagSpec := joinTags(tags...)

	client.Count("events.set_price_batch.size", int64(size), tagSpec, 1)
}
