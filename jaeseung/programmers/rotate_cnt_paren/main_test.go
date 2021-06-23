package main

import (
	"log"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "tc1", args: args{s: "[](){}"}, want: 3},
		{name: "tc2", args: args{s: "}]()[{"}, want: 2},
		{name: "tc3", args: args{s: "[{([[[([[[(({([[[{[([[({{(({{[([([[{{[(({{({((({{({({{(([{({{({(({{[({{{([{[([[{[[[[([{({([((({{[{([{(((({({{([{{[{({{{[[[([[(({{{{((({[[({[{{[{{({{{([{{({{[{([[[({{({[{({([({({[{{[{([[{[([{{[[[(([({{{({({[{[(([{[[[{[{({{[({[[[({[({([((({[([{{[{[([(([[(({[{(([[[{(((([{((([{({([([{([[{{{[{{[{({[({[{{[{((((({{({((({[[[{{([{{[([([([([[[{[(({{({[[({({((((({{({[([{({{{[[({((({[((((([[[({{{[[[({{(((({(({{{(({[[(([([{({([((([{{[{(((([({[[{[[{([{(({[{([([(((((({[{[{{{(([({[({([[[[[{{{(([[[([{{{[{{([{(({(((({{{{[[{({[{{{[([[{({[{{[[{{[{{{(({({{[{(({[{{[([({{({[[{{((([((([{(([([{[[((({{{[{([({({[[[([{((([{{{{[{((([{{[({({[{([[[(((({[((([[{{([{([{[{[([({{{(({(({((([(([{[{[[([(([[{(([{([([[({[[{[{[[[[{{{{{[{{{([(({[[([{{([[({{{[((((({[[[[([([{({({[[(({([[[[{[[({([[{({{([[[{{{({([[{{{[[[[{([{{(([{[(({({([(([[[({[{(({{{(({((([{{([[{([(({[([{{[[{([[{[(([[[(({{{{(([{{[[{[[([({([({[{{({[([{{{(([({[((({{{{[{{[[[{([{(([[{[({(({[{{([({[{{([([(([{[(([{[{[[{{[{{[[[(((([(([[{[(({[[(([{[[[[{{({((({({[({[(([{[(({{([({({({{(([{([[{{(([[([({{([{(({([[[{{{([[[({(((({[[[[[(([[{[({{{[{(({{[({([[({([{([{[[(([{{[({{({[[({{[{(([{({(([[[{{({([([{({[{{{(({[{{{([{{((([[[[[([[{[[[{(((([{{[{{[[([[[{[({{[([[[{[({(((({[{[{[{({[{{[{{[{[([{[({[{([[(([[[[({(({[[[([([{[([((({({{({({{({[{[[{[{[[[{{([([([{{([({([({([{{{{[{(({((([{[{{((({{[([({{{{({[([{[{({[{([[[{(({{[{{([[{{([[[{({[([{[[{{[({{({[[(({{[[[[({[{[[[(([{([{[((([{{({([([(({[(({[[[[([{{[[([{([[(([([([([[((((([[{([{(({{{[{[([{[({{[{[{(({[[([({((({{[[{[{[{(([{[({[((({[[[((([{[{{[((({[[([{{[{({[[[{(({{[{[{([{([({{[[({({{{[{(((([{{{([[[((({{{{[[({((([[({([[{{[([([{{({{((({({({[[{{[{([([{[{{[{[{(({(([[{{{{[[[[({[(([{[{[([[[{[[{{(({[[{(([[[([[(({([({({[{([[[([{[{({{{[{{{{(([{{{({(((((({[{[[(([{[[([([{(({[([{([((([[({{({([{{[{[([({[[([[[{{{({{({([({([({((({{{{[([[{{({({[([{({((({[(([[(([(([({({{(((({{({(({[[(({([{[{([{{{((([(({[({[[({((((([[(({{{{({{([[{[{{{[([{[{(((({(([((([[[(([{(({(({[({{({[{{[([{{{[{[(((((((((({(((({{{{{([[[({([{([[({{[({{{{(([{({([[{((([[{[[[((({(({[{{((([({{[{{([([[({[{{{[([({([{{([{[{({{({[(((([[([([{{{[{{{{{((({[([{{([[{{[({({{[[[[([[[[[(({(((([({{[[{[[(([[[(({{{[[([([(({[[[({([[[{((([{[[(({[{[(((({[[[[{(([[((({{{({[[{[(({(([([{[(([({([[[[({([[{([({{[[{{([{([({({[{[([([{[({[{[{({[(([[{[{([[({({[{((([[([({{{(((([{([[[[{(([{([[{[{[({[{({[(([{{[{[{{[([[({[{{{[{({({(({[{{[([(([{[[{[([([{[[{[[[({[[([[({([({{[[(({[[{{{{[[[{[{([({[({[[([{(([{({{[({([({(([[{(((({((([(([[{[[[([[{[([(({(({{({{{{(([{({{({{{({[[[([[(((([{{{{{([[[({([[{[[({([[[[[[[{({{{{[{([[{{({[{{[[{[{[[[([(({{[[[{{({[[{({({((([((({({{({(([{[([[{({([{{((((([([(([([({{{[([({([{{{{[({{(({({[(([((({(([{([{[{[({[(([(({{{[{(([[{[([{[[[{{[({{{({{{[[{{([{[{(([[{{{([{((({{[[((([(({{{{(({([{(([{{{[[{{({(({({(({(((([({{{{[((({[[({{[[{(({[([([{([(([(({[([({{{{[[[([[[{{{[([[{{[[[[{({([([[(([{((({[{([{((([([({[{(((({{((({(((({[{[[[{{{(([{{[[[[[([([[[([([([(([[({{([{(({{{(([[((({[([[[{(({([([{[[({((({[{(([[([{([[{(([(([[{[[([{{[({{([[[((({[((([({[[({{[(((({{({[[[[{{{({{{{({([[[[[(((([((({{[({[{{(([([[([[[[([{({{{(([{[[[((((([({{({({{((({([{[[[([{({([((([([[({([({[[(([[[{{{{{[{((([{[([({(([[[{[[(({{[{[{[[{{({([({{[[(([{{{({{{({[{{{(({[((([[{{({[{{[(({[({[{[[[(([[[{[([{([{[{[[{(((([[[({{[{{((({({([[([{[((([{[[[((((([[{([[({([[([[(({[({{([((({([[({([{[[([[[(([[{[[(({[(([{(((((({{({([([([{[[[(([{[[(([([([{[{{{((([{{([{({([{[[{([{(([[[{{[([({{({((({({(({{[(({{[{([({[(({({[[[[[[[((({[[([{(([[[({[[{{({{{({([{([{({[{([[[((({((([([{([[(((({[({(([{[([[[([(({[{{{{[{(({[[[(([([([[(({[[({(((({([{[(({{(([{([({[(({((([[{({{(([{{[{({[[{([((((([(({({[{{{{[{[({{{([[{({(([([(([[{(({([(({{{({{[((({({[{{(([({([((({([[[{{{{{[{[[((([[(([{[[(([[({[[[[[(({[(({{{[[[{{{({[[([{([(((([({{{{({({[{((({[{[[[([[({({([[{{[([[([{([({({{{{{[{[[{[{[{{[(([[{(([[[[(([(({{({{({({{({([[(([{({(([{([(({{{[[[[[(([([{(({[({(((({{[{{[[[(([{{(([[({({{[([[[(({[[([(((({([({[([[({{[{{{(({[{[[([[[([[[(({[((({({{[((({{[[({[{[[{(([[([([[[{[{({{([[[{([{[{{{{({[(([[{[{({[([{((({(((([([[(((((([([[[{{[[(([([([[({([{(({[[[{(({[{({([[(({{({({[{{({[(((({({{[([{{(({({{([[[[([({[{[[({{([({[{[{([((({{[{{{{{{[{[{({([({[[[{[(({{({[([[{[[({[[{[{[{[[(([({({({[{(({(({[(((({(({{{{{[[[{[(({[[((({({([{[{({[{(({[[[({([[[((({{([{([[[({({([(([[{[{{({{([{{{[[{([{{((({{[{[{{(({{(({(({((([[([[{{[{[({{({{{[{{[{([({[{{[([[{({({{{(([[(([{[((({[([({([{{[{[(([[[[{[{{{([[[({{{[{((((([[{{[{([{{[{{(([[([{{[[([[{([[{[[{{[([{{{{[{[([((([((([([{[({{({(([[[({({[(([{[({({{{[{{{{{{(({{({{({{(([[([({[[{{{({[(({{{([[{({{[[([[{[[[[{{([[{(([({[[[[({(({{{[([([(({[(((({((({(([{[((([[({[({([{[([[[[[[(({(((({{{(({[[({((((({([{{{{(([((({[([[[(({[[[{{[([([([({[{(([[({{{[[{((({[[[([{[[({([[[([[[{[([[{([[[({{{{{[[{[{[{[{{((([{{{{{{({{{{{({[[({[({[((({[{([[{{[([{[{{((([([{([([{([{[[{({[({((((([({([([({(({[{({([({{(({({{{([({{[({[({{{({(([[[({{{([({(({[(([([[([{(({{[[({({((({{[[[[{({((((([{[[{[{(({[{[[[([(([([(({({[[([({[[({([([({({{(({{[(([{{[{[([({(([(([[[[[[[[{({{(([([[{[(({({{([{[({[({{[[{({{{{{{((([(({(({[[(({{[{({[{({[{([{[{{{[[[]]]}}}]}])}]})}]})}]}}))]]}))}))])))}}}}}})}]]}})]})]}])}})}))]}]])]))}})}]]]]]]]]))]))})])]}]}}]))]}}))}})})])])})]]})])]]})}))])]))])]]]}]}))}]}]]}])))))})}]]]]}})))})})]]}}))}])]])]))]}))})])}}})]]]))})}}})]})]}})])}}})}))}})])})}]}))})])])})])))))})]})}]]}])}])])}])])))}}]}])]}}]])}]})))]})]})]]})}}}}})}}}}}}])))}}]}]}]}]]}}}}})]]])}]])]}]]])]]])})]]}])]]]})))}]]}}})]]))}]})])])])]}}]]]}))]]])]})))]))}}}}])})))))})]]}))}}}))))}))]]]]]])]}])})]})]])))]}]))})))}))))]}))])])]}}}))})]]]]})]))}]])}}]]]]}]])]]}})}]])}}}))]})}}}]]})])]]))}})}})}}))}}}}}}]}}})})]}]))]})})]]]))})}})]}])])))])))])]}]}}}}])]}}]]}]])}]])]]}}])]]))}}]}}])}]}}]])))))}]}}})]]])}}}]}]]]]))]}]}}])})])]})))]}]))]]))}}})})}]])]}}]})])}]}}]}}})}})]}]}}]])]])))}))}))}}))}}]}]}})))}}])}]]}}}])}})}}]}]]))])})})]]])}])}})))]]])})]]]}))}]})}]}])})})))]]}))]}]]]}}}}}))}))))]}))}))}]})})})]))]]}]}]}]]})]]}]])]})}}))]}]]]})])})}]}]}}}}}}]}})))])}]}]})])}})]]}]})])]]]])}})}))}}])]}})}))))]})}}]})})}}))]])})}]}))}]]]}))}])})]])])]))]]}}]]])]))))))]])]))))})))}])]})}]}]]))]})}}}}]}])}]]])}})}]}]]])])]]))}]]}]})]]}})))]}})})))]}))]]])]]])]]}]}))}}}]}})]])]})])}))))])]]}))]]])]}})})]]))}}]))]]]}}]}}))))})]}))}])]))]]]]]}}}))])}]))})}]))]])})}})})}})}}))]))]]]]))}]]))]}}]}]}]]}]}}}}})})])}])]])]}}]])})})]])]]]}]})))}]})})}}}})]))))])}])]]})}}}]]]}}}))]}))]]]]]})]]))]]}]))]])))]]}]}}}}}]]])})))])})]))}}]})})))]}})}}}))])}))}]]))])]))})}]])}}})]}]}}}}]})}))])))))])}]]})}]}}]))}})}]])))}))]})])}]))}}))]}])}))))})]]}))]])])]))]]]}))}]}}}}]}))])]]])]}]))})]}))))]])}])])))})))]]])}]})}])}])})}}})}}]]})]]]))}])]]})))]]]]]]]})}))]})])}]}}))]}}))})})))})}})])]}}]]]))}])}]]}])})}])}}])))}}}]}])])]))]]}]))]]]}])])])})}}))))))}]))]}))]]}]]))]]])]]}])})]])})))])}})]}))]])]])})]])}]])))))]]]}])))]}])]])})})))}}]}})]]]))))}]]}]}])}])]}]]]))]]]}]})]}))]}}]})}}]])))]}))}}}]})}}})}}}]))]]}})])})}}]]}]}]}}))]]}]]]))})])]}])))}]}}}}}]]]))]]})])})]])])))])})}])]]]}])})))}})})}})])))))]]]}]))}}})}])]]]])]])]))}}]})]}})))]))))]]]]])})}}}})}}}]]]]})}}))))]}})]]})])))]})))]]])}})]}}])]]}]]))]))}]])}])]]))}]})))})]]}])])}))}]]])]})))]]))}}}))}])}})]]))])])])]]])])]]]]]}}]))}}}]]]}]}))))})))}}))))}]})])])))}])}]})))}]))]])])})}]]]]}}]])]}}}]]])]]]}}}})])]}))]))])}])])]}))}]]}})]]})))]}}}})]))))}))})}))})}}]]}}}]))}])}))}}}}))])))]]}})))}])}}}]]))}]}])}}]]}}})}}})]}}]]]}])]}]]))}]}}}))]))]})]}]}])}]))})))]))]})}))}})]}}}}])})])]}}})])]))])])))))}}])})}]])]}]))})}})})))])))})})}]]})}}]]]}}))])]]]}]}]]}}]})}}]])}]}}}})}]]]]]]])})]]}]])})]]])}}}}}]))))]])]]]})}}})}})}]))}}}})}}))}))])]}]])]]]}]]))])))}))))}]]))})])})]}})}]))}])]]})]})])}]}]]]}}}}]]}))]]}})])})]])]]})]]]}]]}])])]}]]}]))])]}}]}))})})}]}}}]})]])]}}]}]}}]))]})}]})]}]}]])}]))}]]]])}]))))}}})])]])))}]})})]])}]}]]))]})}]}]})]}])])]}]})})])}])}}]]}})])}]])})]]]])})]))]}])]))}))]}]]})}}})))]]))}]]]]}))))]}]}))]]}])))}]]])})]]]}))])])]]}}}))]]]))]]}]]}})]))))}))]]]]])]]]]}})})]}}]])}}])]})))}}}}}]}}}])])]]))))]})}})}]}])}}])})])]}}}]})]])])}}]}})])))}}]}))})))]]]}]])))}]])})}]))}}}})]}})]])}])})]]])}}}}}))))}))))))))))]}]}}}])]}}]})}})]}))}))}]))]]])))]))}))))}]}])]}}}]}]])}})}}}}))]])))))})]]})]}))])))}}}])}]}])}))]]}))})}}))))}})})]))]))]]))]})))})}])]})})}}]])]}}}})))})])})])})}})}}}]]])]]})])]}]}}])})}})]])))])}])]}))}])])]]}]))]]}]}))))))})}}}]))}}}}]}}})}]}])]]])}]})})])}))]])]]]))}]]}))}}]]}]]])]}]}]))]})]]]]}}}}]]))}))}]}]}}]}])])}]}}]]})})})))}})}}])])]}}]])})]])))})]]}}}})))]]])}}}]))))}]}}})})]]}})])}])}]}]}}))}]]]})}]}}])]]})))]}}]}])))]]]})))]})]}]))}]}]}]]}})))})])]]}))}]}]}})]}])]}]}}}))}])}]])))))]])])])]))]])}])]]}}])]]]]}))]}))])])})}}])))]}])}]))]]]}]})]]]]}}))]]})}})]}}]]}])]})}]]])}}]])}}]}}))}]]])}]})}]}])]})}}}})])]}})))}}]}])))}))}]}}}}])})])})])}}])])])}}]]]}]}]]}]})}})})}})})))])]}])])]]]}))})]]]]))]])}]})]}])]}]}}]}}]})}]}]}]}))))})]}]]])]}})]}]]])]]}}]}}]))))}]]]}]])]]]]])))}}])}}}]}))}}}]})}])])})}}]]]))})}]))}]}})]]})}})]}}]))]]}])}])})]])})]}}))}]}}})]}]]))]]]]]}))))})]]])}}}]]])}))}])}})])]]))}}]])}]))}})})})])}}))]}]))]})]})})))})}}]]]]}]))]]}))]}]]))]))))]]]}}]}}]]}]}]))]}]))])])}}]})])}}]}))})]}]]))}])}]]]}}]}}}})))]})]))}}}])]})}}]})])})])]]}]]}}]))}}}}))]]]))]}]])}]]}}])]}))])}]])}}])))}))}}}))}]})]]]))])})}))]}]))}}])}]]]]}}}]])})}}}]]])}})}]])})]]}]]]])}))]]})})}])])]]]]})))))]}}})]])}}])]]}))])}}}]}}}}}]]]]}]}]]})]])])}]))}]]))])]]}]}]))])))}))}))}}})])]}]}])}])}}]])))]}))))]]])}]})})]}}])))}]}}}}])))}])]]]})})])}]}}})))]]}])]))}])))])))}}]]})}})])]}}]}))}]}})}))}}}]}}]]}}]})}]])]}}}]})}]]}}}}))))}))}])}}]}}}])]]]))}}}]]]]])})]})]))}}}]}]}))))))])])}]}))}])}]]}]]})]))))}]}}])))])})}])]))]]}))}}}))}))))}})]]]}}})]]])))))]})))})]]}}})}])]})}})))))})})]]})}}))]}]]])])])])]}}])}}]]]})))})}})))))}]}}]})]})}]}}]}}}]])}])])})}])))}]))))}]]]))}]}))]]))])]}]}}])]})))])})]})]]]})]}})}]}]]]}]))]}]})})}}})]))]]]}}])]}]])}]}}]})})])})}]})}})]]])}]}})}}])}}})}}]}}]})]]})))}}}}))]])]]]}}})}]}}])}})}))))}])}]}})))])})}])]]]]}]])]}])}}})]}}))})}})}]))}})})}})))})}}))]}}]])])]}}))}})]])]}]]])}))]]])]]])}]"}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.s); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTest(t *testing.T) {
	m := map[string]int{}
	var mm map[string]int
	log.Println(m)
	log.Println(mm)
}