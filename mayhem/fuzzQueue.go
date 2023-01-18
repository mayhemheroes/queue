package fuzzQueue

import "strconv"
import "github.com/smallnest/queue"


func mayhemit(bytes []byte) int {

    var num int
    if len(bytes) > 1 {
        num, _ = strconv.Atoi(string(bytes[0]))

        switch num {
    
        case 0:
            test := queue.NewLKQueue[int]()
            for byte := range bytes {

                if byte % 2 == 0 {

                    test.Enqueue(byte)
                } else {

                    _ = test.Dequeue()
                }
            }
            return 0

        case 1:
            test := queue.NewSliceQueue[int](64)
            for byte := range bytes {

                if byte % 2 == 0 {

                    test.Enqueue(byte)
                } else {

                    _ = test.Dequeue()
                }
            }
            return 0

        case 2:
            test := queue.NewBoundedQueue[int](uint32(64))
            for byte := range bytes {

                if byte % 2 == 0 {

                    test.Enqueue(byte)
                } else {

                    _ = test.Dequeue()
                }
            }
            return 0

        default:
            test := queue.NewCQueue[int]()
            for byte := range bytes {

                if byte % 2 == 0 {

                    test.Enqueue(byte)
                } else {

                    _ = test.Dequeue()
                }
            }
            return 0

        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}