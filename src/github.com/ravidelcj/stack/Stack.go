package stack


//Data type that defines one identity
type element struct {
  title string
  date string
  url string
  remoteUrl string
}

type StackNode struct {
  e []element
}
//append data to the stack
func (st *StackNode) push (data element) bool {
    st.e = append(st.e, data)
    return true
}

//pop from stack
func (st *StackNode) pop () bool {

   if len(st.e) == 0 {
     return false
   }

   //deleting last element
   st.e = st.e[ :len(st.e) - 1 ]
   return true
}

//check whether the stack is empty or not
func (st *StackNode) isEmpty() bool {
  if len(st.e) == 0 {
    return true
  }
  return false
}

func (st *StackNode) size() int {

  return len(st.e)

}

func (st *StackNode) top() (element, bool) {
  if st.isEmpty() == true {
    return element{title : "title", date : "date", url : "url", remoteUrl : "remoteUrl"},false
  }
  return st.e[len(st.e)-1], true
}
