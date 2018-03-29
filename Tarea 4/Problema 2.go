package main
//import "fmt"
type BloomFilter struct {
    bit [100]bool
    f1 func (int) int
    f2 func (int) int
    f3 func (int) int
    f4 func (int) int
}
func (B *BloomFilter) insert(A []int) {
    for i := 0; i < len(A); i++ {
        B.bit[B.f1(A[i])] = true
        B.bit[B.f2(A[i])] = true
        B.bit[B.f3(A[i])] = true
        B.bit[B.f4(A[i])] = true
} }
func (B *BloomFilter) member(x int) bool {
i:=0
j:=true
if B.bit[B.f1(x)]==true{
  i++
}
if B.bit[B.f2(x)]==true{
  i++
}
if B.bit[B.f3(x)]==true{
  i++
}
if B.bit[B.f4(x)]==true{
  i++
}
if i!=4{
  j=false
  return j
}
return j
}
func main() {
    var B BloomFilter
    B.f1 = func (x int) int
        { return ((42727 * x + 4047) % 33613) % 100 }
    B.f2 = func (x int) int
        { return ((12562 * x + 57701) % 59771) % 100 }
    B.f3 = func (x int) int
        { return ((40714 * x + 25926) % 29101) % 100 }
    B.f4 = func (x int) int
        { return ((48752 * x + 5084) % 7193) % 100 }
    A := []int{522, 1641, 3295, 4965, 7853, 8544, 9256, 9723}
    B.insert(A)

}
