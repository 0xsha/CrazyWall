package modules

//type DomainInfo struct {
//	Awis struct {
//		OperationRequest struct {
//			RequestID string `json:"RequestId"`
//		} `json:"OperationRequest"`
//		Results struct {
//			ResponseStatus struct {
//				StatusCode string `json:"StatusCode"`
//			} `json:"ResponseStatus"`
//			Result struct {
//				Alexa struct {
//					ContentData struct {
//						AdultContent string `json:"AdultContent"`
//						DataURL      string `json:"DataUrl"`
//						Language     struct {
//							Locale string `json:"Locale"`
//						} `json:"Language"`
//						LinksInCount string `json:"LinksInCount"`
//						SiteData     struct {
//							Description string `json:"Description"`
//							Title       string `json:"Title"`
//						} `json:"SiteData"`
//						Speed struct {
//							MedianLoadTime string `json:"MedianLoadTime"`
//							Percentile     string `json:"Percentile"`
//						} `json:"Speed"`
//					} `json:"ContentData"`
//					Request struct {
//						Arguments struct {
//							Argument []struct {
//								Name  string `json:"Name"`
//								Value string `json:"Value"`
//							} `json:"Argument"`
//						} `json:"Arguments"`
//					} `json:"Request"`
//					TrafficData struct {
//						ContributingSubdomains struct {
//							ContributingSubdomain []struct {
//								DataURL   string `json:"DataUrl"`
//								PageViews struct {
//									PerUser    string `json:"PerUser"`
//									Percentage string `json:"Percentage"`
//								} `json:"PageViews"`
//								Reach struct {
//									Percentage string `json:"Percentage"`
//								} `json:"Reach"`
//								TimeRange struct {
//									Months string `json:"Months"`
//								} `json:"TimeRange"`
//							} `json:"ContributingSubdomain"`
//						} `json:"ContributingSubdomains"`
//						DataURL       string `json:"DataUrl"`
//						Rank          string `json:"Rank"`
//						RankByCountry struct {
//							Country []struct {
//								Code         string `json:"@Code"`
//								Contribution struct {
//									PageViews string `json:"PageViews"`
//									Users     string `json:"Users"`
//								} `json:"Contribution"`
//								Rank string `json:"Rank"`
//							} `json:"Country"`
//						} `json:"RankByCountry"`
//						UsageStatistics struct {
//							UsageStatistic []struct {
//								PageViews struct {
//									PerMillion struct {
//										Delta string `json:"Delta"`
//										Value string `json:"Value"`
//									} `json:"PerMillion"`
//									PerUser struct {
//										Delta string `json:"Delta"`
//										Value string `json:"Value"`
//									} `json:"PerUser"`
//									Rank struct {
//										Delta string `json:"Delta"`
//										Value string `json:"Value"`
//									} `json:"Rank"`
//								} `json:"PageViews"`
//								Rank struct {
//									Delta string `json:"Delta"`
//									Value string `json:"Value"`
//								} `json:"Rank"`
//								Reach struct {
//									PerMillion struct {
//										Delta string `json:"Delta"`
//										Value string `json:"Value"`
//									} `json:"PerMillion"`
//									Rank struct {
//										Delta string `json:"Delta"`
//										Value string `json:"Value"`
//									} `json:"Rank"`
//								} `json:"Reach"`
//								TimeRange struct {
//									Days   string `json:"Days"`
//									Months string `json:"Months"`
//								} `json:"TimeRange"`
//							} `json:"UsageStatistic"`
//						} `json:"UsageStatistics"`
//					} `json:"TrafficData"`
//				} `json:"Alexa"`
//			} `json:"Result"`
//		} `json:"Results"`
//	} `json:"Awis"`
//}
//
//func GetDomainInfo()  {
//
//}
//
//
//type TrafficHistory struct {
//	Awis struct {
//		OperationRequest struct {
//			RequestID string `json:"RequestId"`
//		} `json:"OperationRequest"`
//		Results struct {
//			ResponseStatus struct {
//				StatusCode string `json:"StatusCode"`
//			} `json:"ResponseStatus"`
//			Result struct {
//				Alexa struct {
//					Request struct {
//						Arguments struct {
//							Argument []struct {
//								Name  string `json:"Name"`
//								Value string `json:"Value"`
//							} `json:"Argument"`
//						} `json:"Arguments"`
//					} `json:"Request"`
//					TrafficHistory struct {
//						HistoricalData struct {
//							Data []struct {
//								Date      string `json:"Date"`
//								PageViews struct {
//									PerMillion string `json:"PerMillion"`
//									PerUser    string `json:"PerUser"`
//								} `json:"PageViews"`
//								Rank  string `json:"Rank"`
//								Reach struct {
//									PerMillion string `json:"PerMillion"`
//								} `json:"Reach"`
//							} `json:"Data"`
//						} `json:"HistoricalData"`
//						Range string `json:"Range"`
//						Site  string `json:"Site"`
//						Start string `json:"Start"`
//					} `json:"TrafficHistory"`
//				} `json:"Alexa"`
//			} `json:"Result"`
//		} `json:"Results"`
//	} `json:"Awis"`
//}
