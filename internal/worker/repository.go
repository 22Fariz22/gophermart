package worker

type WorkerRepository interface {
	SendToOrdersCannels()
	SendToAccrualBox(orders []NewOrders) error
	SendToWaitListChannels()
}
