package main

import (
	"encoding/json"
	"fmt"
)

type (
	FlowConfig struct {
		CapSetStrategy  string                 `json:"set_exceeded_strategy"`
		ExternalClients []ExternalClientConfig `json:"external_clients"`
		PreCapData      []ValueGenerator       `json:"pre_cap_data"`
		PostCapFilter   []PostCapFilter        `json:"post_cap_filter"`
	}

	ExternalClientConfig struct {
		ClientName       string           `json:"client_name"`
		ClientParameters []ValueGenerator `json:"client_parameters"`
		QueryClientWhen  WhenExpression   `json:"query_client_when"`
	}

	PostCapFilter struct {
		FilterName string         `json:"filter_name"`
		FilterWhen WhenExpression `json:"filter_when"`
	}

	ValueGenerator struct {
		Name  string      `json:"name"`
		Value ValueFinder `json:"value"`
	}

	WhenExpression struct {
		Operation   string       `json:"when_operation"`
		Expressions []Expression `json:"expressions"`
	}

	ValueFinder struct {
		FunctionValue FunctionValue `json:"function_value"`
		SingleValue   string        `json:"single_value"`
	}

	Expression struct {
		Operation string      `json:"operation"`
		Left      ValueFinder `json:"left"`
		Right     ValueFinder `json:"right"`
	}

	FunctionValue struct {
		Function   string            `json:"function"`
		Parameters FunctionParameter `json:"parameters"`
	}

	FunctionParameter struct {
		Origin []string `json:"origin"`
		Field  []string `json:"field"`
	}
)

func main() {

	var entryString string
	entryString = noScapeStringFunction(false)
	dataNoScape := []byte(entryString)

	var flowConfig FlowConfig
	err := json.Unmarshal(dataNoScape, &flowConfig)

	if err != nil {
		fmt.Println("Error")
	}

	//fmt.Printf("%+v\n",flowConfig)
	fmt.Printf("%+v\n", flowConfig.ExternalClients[0].QueryClientWhen.Expressions[0].Right.FunctionValue.Function)
	fmt.Printf("%+v\n", flowConfig.ExternalClients[0].QueryClientWhen.Expressions[0].Right.FunctionValue.Function == "")
}

func noScapeStringFunction(inline bool) string {
	if inline {
		return ""
	}
	return `
	{
		"set_exceeded_strategy" : "single_cap",  
		"external_clients": [
			{
				"client_name": "datavault",
				"client_parameters" : [
					{
						"name": "user_id",
						"value": {
							"function_value": {
								"function": "GET",
								"parameters": {
									"origin": ["FlowParamater"],
									"field": ["receptor_identification_number"]
								}
							}
						}
					}
				],
				
				"query_client_when": {
					"when_operation": "AND",
					"expressions": [
						{
							"operation": "NE",
							"left": {
								"function_value": {
									"function": "GET",
									"parameters": {
										"origin": ["FlowParamater"],
										"field": ["receptor_identification_number"]
									}
								}
							},
							"right":  {
								"single_value": ""
							} 
						}
					]
				}
			},
			{
				"client_name": "trusted_contacts",    
				"client_parameters" : [
					{
						"name": "customer_id",
						"value": {
							"function_value": {
								"function": "GET",
								"parameters": {
									"origin": ["GeneralParameter"],
									"field": ["user_id"]
								}
							}
						}
					},
					{
						"name": "user_id",
						"value": {
							"function_value": {
								"function": "GET",
								"parameters": {
									"origin": ["FlowParameter","FlowParameter"],
									"field": ["receptor_id", "receptor_identification_number"]
								}
							}
						}
					},
					{
						"name": "identification_type",
						"value": {
							"function_value": {
								"function": "GET",
								"parameters": {
									"origin": ["ExternalResponse-datavault","SingleValue"],
									"field": ["User.Identification.Type", "receiver_user_id"]
								}
							}
						}
					}
				],    
				"query_client_when": {
					"when_operation": "ALL",
					"expressions": []
				}
			}
		],
		"pre_cap_data": [
			{
				"name": "person_type",
				"value": {
					"function_value": {
						"function": "GET",
						"parameters": {
							"origin": ["ExternalResponse-datavault","ExternalResponse-trusted_contacts", "SingleValue" ],
							"field": ["User.Identification.Type", "Identification.Type","CPF"]
						}
					}
				}
			},
			{
				"name": "transaction_type",
				"value": {
					"function_value": {
						"function": "GET",
						"parameters": {
							"origin": ["FlowParam"],
							"field": ["transaction_type"]
						}
					}
				}
			},
			{
				"name": "trusted_contact",
				"value": {
					"function_value": {
						"function": "GET",
						"parameters": {
							"origin": ["ExternalResponse-trusted_contacts","SingleValue"],
							"field": ["trusted_contact","false"]
						}
					}
				}   
			}
		],
		"post_cap_filter": [
			{
				"filter_name": "post_time_slot",
				"filter_when": {
					"when_operation": "AND",
					"expressions": [
						{
							"operation": "GTE",
							"left": {
								"function_value": {
									"function": "GET",
									"parameters": {
										"origin": ["CapsResponse"],
										"field": ["Cap.LimitDates.Gte"]
									}
								}
							},
							"right":  {
								"function_value": {
									"function": "NOW",
									"parameters": {
										"origin": ["FlowParamater"],
										"field": ["site_id"]
									}
								}
							} 
						},
						{
							"operation": "LE",
							"left": {
								"function_value": {
									"function": "GET",
									"parameters": {
										"origin": ["CapsResponse"],
										"field": ["Cap.LimitDates.Lte"]
									}
								}
							},
							"right":  {
								"function_value": {
									"function": "NOW",
									"parameters": {
										"origin": ["FlowParamater"],
										"field": ["site_id"]
									}
								}
							} 
						}
					]
				}
			}
		]
	}
	`
}
