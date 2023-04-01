package main

import (
	"gopkg.in/yaml.v2"
	"fmt"
)

type (
	FlowConfig struct {
		CapSetStrategy  string                 `yaml:"set_exceeded_strategy"`
		ExternalClients []ExternalClientConfig `yaml:"external_clients"`
		PreCapData      []ValueGenerator       `yaml:"pre_cap_data"`
		PostCapFilter   []PostCapFilter        `yaml:"post_cap_filter"`
	}

	ExternalClientConfig struct {
		ClientName       string           `yaml:"client_name"`
		ClientParameters []ValueGenerator `yaml:"client_parameters"`
		QueryClientWhen  WhenExpression   `yaml:"query_client_when"`
	}

	PostCapFilter struct {
		FilterName string         `yaml:"filter_name"`
		FilterWhen WhenExpression `yaml:"filter_when"`
	}

	ValueGenerator struct {
		Name  string      `yaml:"name"`
		Value ValueFinder `yaml:"value"`
	}

	WhenExpression struct {
		Operation   string       `yaml:"when_operation"`
		Expressions []Expression `yaml:"expressions"`
	}

	ValueFinder struct {
		FunctionValue FunctionValue `yaml:"function_value"`
		SingleValue   string        `yaml:"single_value"`
	}

	Expression struct {
		Operation string        `yaml:"operation"`
		Left      ValueFinder `yaml:"left"`
		Right     ValueFinder `yaml:"right"`
	}

	FunctionValue struct {
		Function   string            `yaml:"function"`
		Parameters FunctionParameter `yaml:"parameters"`
	}

	FunctionParameter struct {
		Origin []string `yaml:"origin"`
		Field  []string `yaml:"field"`
	}
)

func main() {
	var entryString string
	entryString = noScapeStringFunction()
	dataNoScape := []byte(entryString)

	var flowConfig FlowConfig
	err := yaml.Unmarshal(dataNoScape, &flowConfig)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("%+v\n",flowConfig)
}

func noScapeStringFunction() string { 
	return "{set_exceeded_strategy: single_cap, external_clients: [{client_name: datavault, client_parameters: [{name: user_id, value: {function_value: {function: GET, parameters: {origin: [FlowParamater], field: [receptor_identification_number]}}}}], query_client_when: {when_operation: AND, expressions: [{operation: NE, left: {function_value: {function: GET, parameters: {origin: [FlowParamater], field: [receptor_identification_number]}}}, right: {single_value: ""}}]}}, {client_name: trusted_contacts, client_parameters: [{name: customer_id, value: {function_value: {function: GET, parameters: {origin: [GeneralParameter], field: [user_id]}}}}, {name: user_id, value: {function_value: {function: GET, parameters: {origin: [FlowParameter, FlowParameter], field: [receptor_id, receptor_identification_number]}}}}, {name: identification_type, value: {function_value: {function: GET, parameters: {origin: [ExternalResponse-datavault, SingleValue], field: [User.Identification.Type, receiver_user_id]}}}}], query_client_when: {when_operation: ALL, expressions: []}}], pre_cap_data: [{name: person_type, value: {function_value: {function: GET, parameters: {origin: [ExternalResponse-datavault, ExternalResponse-trusted_contacts, SingleValue], field: [User.Identification.Type, Identification.Type, CPF]}}}}, {name: transaction_type, value: {function_value: {function: GET, parameters: {origin: [FlowParam], field: [transaction_type]}}}}, {name: trusted_contact, value: {function_value: {function: GET, parameters: {origin: [ExternalResponse-trusted_contacts, SingleValue], field: [trusted_contact, 'false']}}}}], post_cap_filter: [{filter_name: post_time_slot, filter_when: {when_operation: AND, expressions: [{operation: GTE, left: {function_value: {function: GET, parameters: {origin: [CapsResponse], field: [Cap.LimitDates.Gte]}}}, right: {function_value: {function: NOW, parameters: {origin: [FlowParamater], field: [site_id]}}}}, {operation: LE, left: {function_value: {function: GET, parameters: {origin: [CapsResponse], field: [Cap.LimitDates.Lte]}}}, right: {function_value: {function: NOW, parameters: {origin: [FlowParamater], field: [site_id]}}}}]}}]}"
}
