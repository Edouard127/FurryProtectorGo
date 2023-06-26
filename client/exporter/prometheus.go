package exporter

import "github.com/prometheus/client_golang/prometheus"

var (
	InterationHist = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "discord_interaction_duration",
		Help: "The number of interaction received by guild per user",
	}, []string{"guild", "user", "interaction"})

	MemberGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "discord_member_count",
		Help: "The number of members per guild",
	}, []string{"guild"})

	MemberJoinCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "discord_member_join_count",
		Help: "The number of member join per guild",
	}, []string{"guild"})

	MemberDeleteCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "discord_member_delete_count",
		Help: "The number of member leave per guild",
	}, []string{"guild"})

	MessageCreateCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "discord_message_number",
		Help: "The number of messages received by guild by channel per user",
	}, []string{"guild", "channel", "user"})

	MessageDeleteCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "discord_message_delete_number",
		Help: "The number of messages deleted by guild by channel",
	}, []string{"guild", "channel"})

	MessageUpdateCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "discord_message_update_number",
		Help: "The number of messages updated by guild by channel per user",
	}, []string{"guild", "channel", "user"})
)

func DoRegister(registry *prometheus.Registry) *prometheus.Registry {
	registry.MustRegister(InterationHist,
		MemberGauge,
		MemberJoinCounter,
		MemberDeleteCounter,
		MessageCreateCounter,
		MessageDeleteCounter,
		MessageUpdateCounter)
	return registry
}
