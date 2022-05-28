const { Kafka } = require('kafkajs')

const kafka = new Kafka({
    clientId: 'my-app',
    brokers: ['localhost:9092']
})

const consumer = kafka.consumer({ groupId: 'test-group' })
let consumer_running = false


async function kafkaSetState(fnSetState) {
    if (consumer_running) {
        return;
    }
    console.log("This ever called?")
    consumer_running = true;

    consumer.connect()
    consumer.subscribe({ topic: 'cpu', fromBeginning: true })
    await consumer.run({
        eachMessage: async ({ topic, partition, message }) => {
            console.log({
                value: message.value.toString(),
            })
            // let json_values = JSON.parse(message.value.toString())
            // json_values["received_first_message"] = true

            // fnSetState(json_values)
        },
    })
}

// kafkaSetState(null)
export default kafkaSetState

