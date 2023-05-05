package cmd

// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/sqs@v1.20.9/types#Message
// FIXME: is there a way I can slurp this instead of copy-paste?

type MessageAttributeValue struct {

	// Amazon SQS supports the following logical data types: String , Number , and
	// Binary . For the Number data type, you must use StringValue . You can also
	// append custom labels. For more information, see Amazon SQS Message Attributes (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-message-metadata.html#sqs-message-attributes)
	// in the Amazon SQS Developer Guide.
	//
	// This member is required.
	DataType *string

	// Not implemented. Reserved for future use.
	BinaryListValues [][]byte

	// Binary type attributes can store any binary data, such as compressed data,
	// encrypted data, or images.
	BinaryValue []byte

	// Not implemented. Reserved for future use.
	StringListValues []string

	// Strings are Unicode with UTF-8 binary encoding. For a list of code values, see
	// ASCII Printable Characters (http://en.wikipedia.org/wiki/ASCII#ASCII_printable_characters)
	// .
	StringValue *string
	// contains filtered or unexported fields
}

type Message struct {
	// A map of the attributes requested in ReceiveMessage to their respective values.
	// Supported attributes:
	//   - ApproximateReceiveCount
	//   - ApproximateFirstReceiveTimestamp
	//   - MessageDeduplicationId
	//   - MessageGroupId
	//   - SenderId
	//   - SentTimestamp
	//   - SequenceNumber
	// ApproximateFirstReceiveTimestamp and SentTimestamp are each returned as an
	// integer representing the epoch time (http://en.wikipedia.org/wiki/Unix_time) in
	// milliseconds.
	Attributes map[string]string

	// The message's contents (not URL-encoded).
	Body *string

	// An MD5 digest of the non-URL-encoded message body string.
	MD5OfBody *string

	// An MD5 digest of the non-URL-encoded message attribute string. You can use this
	// attribute to verify that Amazon SQS received the message correctly. Amazon SQS
	// URL-decodes the message before creating the MD5 digest. For information about
	// MD5, see RFC1321 (https://www.ietf.org/rfc/rfc1321.txt) .
	MD5OfMessageAttributes *string

	// Each message attribute consists of a Name , Type , and Value . For more
	// information, see Amazon SQS message attributes (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-message-metadata.html#sqs-message-attributes)
	// in the Amazon SQS Developer Guide.
	MessageAttributes map[string]MessageAttributeValue

	// A unique identifier for the message. A MessageId is considered unique across all
	// Amazon Web Services accounts for an extended period of time.
	MessageId *string

	// An identifier associated with the act of receiving the message. A new receipt
	// handle is returned every time you receive a message. When deleting a message,
	// you provide the last received receipt handle to delete the message.
	ReceiptHandle *string
	// contains filtered or unexported fields
}
