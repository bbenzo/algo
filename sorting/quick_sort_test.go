package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkQuickSort(b *testing.B) {

	given := []int{57,-24,-133,5,99,-132,-122,264,393,-35,525,-203,310,-47,-170,36,315,-61,-617,-112,-387,-956,-318,-200,-157,-603,261,287,-588,-9,457,409,-544,-357,319,187,5,517,669,-363,166,631,2,-142,-816,-40,680,-581,-495,-528,933,-398,644,-468,-533,-170,7,738,89,337,-199,-308,-200,-790,117,65,-499,68,301,-55,319,-624,-195,-473,-769,-183,693,-584,218,-198,-627,-760,24,-237,-1,-186,24,313,718,380,-797,-737,118,-253,-102,170,41,-372,-64,-166,-913,-265,-582,-423,58,-272,-270,-231,538,683,-514,43,514,-11,-233,-193,349,84,-228,183,277,-32,-500,-846,87,-402,-325,-70,25,323,-274,243,-286,-486,8,656,-26,-90,102,415,107,-98,-170,-409,-14,662,170,-292,291,-222,191,851,152,-676,227,764,-248,324,-283,246,-65,197,-506,-100,298,234,595,809,-28,-353,-656,628,187,278,-97,519,384,-371,-284,-36,293,452,370,-152,182,562,-78,257,-388,285,790,-584,-54,263,3,559,-820,9,-426,384,228,426,-9,-347,-383,-235,825,-370,558,255,405,685,170,-348,363,139,39,881,-104,223,-67,430,232,863,-581,594,142,226,-624,-231,271,439,5,-388,746,-279,-622,-455,-427,-431,596,-619,-521,155,430,-396,-304,157,677,583,-615,-334,-297,490,110,904,493,-887,-475,-75,-24,-439,635,77,179,-619,60,-842,-76,379,281,8,752,942,318,217,-24,868,371,195,-10,-80,524,-16,40,-157,342,681,-465,104,273,-83,-86,-516,418,-81,-290,362,-300,639,-346,-357,-408,-190,30,-76,102,-729,-261,428,132,-149,590,-792,-567,-448,709,-338,-451,-191,914,-735,244,354,145,756,70,-3,418,-51,-339,115,-105,-633,-138,-207,-671,-284,-135,-93,721,665,561,-640,-7,-104,20,667,-177,-535,130,161,-821,251,93,171,355,-118,-638,-424,-419,-233,-641,-93,442,-62,-128,222,501,-29,-197,-135,-337,37,209,254,203,-199,250,101,-610,-647,311,232,365,650,-159,695,312,-366,585,-95,-805,-106,128,-291,544,522,-36,503}
	for n := 0; n < b.N; n++ {
		QuickSort(given)
	}
}

func TestQuickSortWorstCase(t *testing.T) {
	given := []int{57,-24,-133,5,99,-132,-122,264,393,-35,525,-203,310,-47,-170,36,315,-61,-617,-112,-387,-956,-318,-200,-157,-603,261,287,-588,-9}
	expected := []int{-956, -617, -603, -588, -387, -318, -203, -200, -170, -157, -133, -132, -122, -112, -61, -47, -35, -24, -9, 5, 36, 57, 99, 261, 264, 287, 310, 315, 393, 525}

	result := QuickSort(given)

	assert.Equal(t, expected, result)
}

func TestQuickSort(t *testing.T) {
	given := []int{4, 1, 23, 34, 20, 11, 40}
	expected := []int{1, 4, 11, 20, 23, 34, 40}

	result := QuickSort(given)

	assert.Equal(t, expected, result)
}
