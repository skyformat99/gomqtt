package broker

// -- authentication
// authenticate successfully a client with username and password
// authenticate unsuccessfully a client with username and password
// authenticate errors
// authorize publish
// do not authorize publish
// authorize subscribe
// negate subscription
// failed authentication does not disconnect other client with same clientId

// -- basic
// connect and connack (minimal)
// publish QoS 0
// subscribe QoS 0
// does not die badly on connection error
// unsubscribe
// unsubscribe on disconnect
// disconnect
// retain messages
// closes
// connect without a clientId for MQTT 3.1.1
// disconnect another client with the same clientId
// disconnect if another broker connects the same client
// publish to $SYS/broker/new/clients
// restore QoS 0 subscriptions not clean
// double sub does not double deliver

// -- client pub sub
// publish direct to a single client QoS 0
// publish direct to a single client QoS 1
// offline message support for direct publish
// subscribe a client programmatically
// subscribe a client programmatically multiple topics
// subscribe a client programmatically with full packet

// -- events
// publishes an heartbeat
// does not forward $SYS topics to # subscription
// does not store $SYS topics to QoS 1 # subscription

// -- keepalive
// supports pingreq/pingresp
// supports keep alive disconnections
// supports keep alive disconnections after a pingreq
// disconnect if a connect does not arrive in time

// -- meta
// count connected clients
// call published method
// call published method with client
// emit publish event with client
// emit subscribe event
// emit unsubscribe event
// emit clientDisconnect event
// emits client

// -- others
// do not block after a subscription

// -- qos1
// publish QoS 1
// subscribe QoS 1
// subscribe QoS 0, but publish QoS 1
// restore QoS 1 subscriptions not clean
// remove stored subscriptions if connected with clean=true
// resend publish on non-clean reconnect QoS 1
// do not resend QoS 1 packets at each reconnect
// do not resend QoS 1 packets if reconnect is clean
// do not resend QoS 1 packets at reconnect if puback was received
// deliver QoS 1 retained messages
// deliver QoS 0 retained message with QoS 1 subscription
// remove stored subscriptions after unsubscribe
// upgrade a QoS 0 subscription to QoS 1

// -- qos2
// publish QoS 2
// subscribe QoS 2
// subscribe QoS 0, but publish QoS 2
// restore QoS 2 subscriptions not clean
// resend publish on non-clean reconnect QoS 2
// resend pubrel on non-clean reconnect QoS 2
// publish after disconnection

// -- error
// after an error, outstanding packets are discarded

// -- will
// delivers a will
// delivers old will in case of a crash
// store the will in the persistence
