package stack


//Data type that defines one identity
type Element struct {
  Title string
  Date string
  Url string
  RemoteUrl string
}

type StackNode struct {
  e []Element
}

func NewStack() *StackNode  {
  return &StackNode{ e : make([]Element, 0)}
}
//append data to the stack
func (st *StackNode) Push (data Element) bool {
    st.e = append(st.e, data)
    return true
}

//pop from stack
func (st *StackNode) Pop () bool {

   if len(st.e) == 0 {
     return false
   }

   //deleting last element
   st.e = st.e[ :len(st.e) - 1 ]
   return true
}

//check whether the stack is empty or not
func (st *StackNode) IsEmpty() bool {
  if len(st.e) == 0 {
    return true
  }
  return false
}

func (st *StackNode) Size() int {

  return len(st.e)

}

func (st *StackNode) Top() (Element, bool) {
  if st.IsEmpty() == true {
    return Element{Title : "title", Date : "date", Url : "url", RemoteUrl : "remoteUrl"},false
  }
  return st.e[len(st.e)-1], true
}
