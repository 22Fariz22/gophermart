package worker

type UseCase interface {
	SendToOrdersCannels()
	SendToAccrualBox(orders []NewOrders) error
	SendToWaitListChannels()
}
