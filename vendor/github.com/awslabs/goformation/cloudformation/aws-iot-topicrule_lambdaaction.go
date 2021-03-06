package cloudformation

import (
	"encoding/json"
)

// AWSIoTTopicRule_LambdaAction AWS CloudFormation Resource (AWS::IoT::TopicRule.LambdaAction)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-lambdaaction.html
type AWSIoTTopicRule_LambdaAction struct {

	// FunctionArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-lambdaaction.html#cfn-iot-topicrule-lambdaaction-functionarn
	FunctionArn *Value `json:"FunctionArn,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRule_LambdaAction) AWSCloudFormationType() string {
	return "AWS::IoT::TopicRule.LambdaAction"
}

func (r *AWSIoTTopicRule_LambdaAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
