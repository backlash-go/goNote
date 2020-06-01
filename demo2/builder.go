package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {



	var builder1 strings.Builder
	builder1.WriteString("abc")
	builder1.WriteString("abc")

	fmt.Println(builder1.Len(),builder1.Cap()) //0 0
	builder1.Grow(10)
	fmt.Println(builder1.Len(),builder1.Cap())  // 0 10
	fmt.Println(builder1)

	builder1.Reset()
	builder2 := builder1
	fmt.Println(builder2.String(),builder1.String())
	builder2.Grow(2)
	fmt.Println(builder2.Len(),builder2.Cap())  // 0 10


	reader1 := strings.NewReader(
		"NewReader returns a new Reader reading from s. " +
			"It is similar to bytes.NewBufferString but more efficient and read-only.")
	//example 1:
	//fmt.Printf("The size of reader: %d\n", reader1.Size()) //119
	//fmt.Printf("The len of reader: %d\n", reader1.Len()) //119
	//fmt.Printf("The reading index in reader: %d\n",
	//	reader1.Size()-int64(reader1.Len()))// 0
	//
	//buf1 := make([]byte, 47)
	//n, _ := reader1.Read(buf1)	// 从reader中读出buf1大小的内容 返回读取的字节数
	//fmt.Printf("%d bytes were read. (call Read)\n", n) //47
	//fmt.Printf(" aa The size of reader: %d\n", reader1.Size()) //
	//fmt.Printf("a  aThe len of reader: %d\n", reader1.Len()) // 72
	//fmt.Printf("The reading index in reader: %d\n",
	//	reader1.Size()-int64(reader1.Len())) //47
	//fmt.Printf("buf1:%s\n",buf1) //NewReader returns a new Reader reading from s.

	//fmt.Println("this seek is : %s",reader1.Seek())


	//example 2:
	//buf2 := make([]byte, 47)


	offset2 := int64(17)
	//expectedIndex := reader1.Size() - int64(reader1.Len()) + offset2
	fmt.Printf("Seek with offset %d and whence %d ...\n", offset2, io.SeekCurrent)// 17, 1
	readingIndex, _ := reader1.Seek(offset2, io.SeekCurrent)
	fmt.Printf("The reading index in reader: %d (returned by Seek)\n", readingIndex)//17
	//fmt.Printf("The reading index in reader: %d (computed by me)\n", expectedIndex)// 17
	fmt.Println(reader1)
	//n, _ := reader1.Read(buf2)
	//fmt.Printf("%d bytes were read. (call Read)\n", n) //47
	fmt.Printf("The reading index in reader: %d\n", reader1.Size()-int64(reader1.Len())) //64
	fmt.Printf("The size of reader: %d\n", reader1.Size()) //119
	fmt.Printf("The len of reader: %d\n", reader1.Len()) //55
	fmt.Println(reader1)

/*
   buf2 := make([]byte, 21)
       offset1 := int64(64)
       n, _ = reader1.ReadAt(buf2, offset1)
       fmt.Printf("%d bytes were read. (call ReadAt, offset: %d)\n", n, offset1)
       fmt.Printf("The reading index in reader: %d\n",
           reader1.Size()-int64(reader1.Len()))
       fmt.Printf("buf2:%s\n",buf2)
       fmt.Println(reader1)
       fmt.Println()
       n, _ = reader1.Read(buf2)
       fmt.Println(reader1)

       fmt.Println()
       // 示例3。
       offset2 := int64(17)
       expectedIndex := reader1.Size() - int64(reader1.Len()) + offset2
       fmt.Printf("Seek with offset %d and whence %d ...\n", offset2, io.SeekCurrent)
       readingIndex, _ := reader1.Seek(offset2, io.SeekCurrent)
       fmt.Printf("The reading index in reader: %d (returned by Seek)\n", readingIndex)
       fmt.Printf("The reading index in reader: %d (computed by me)\n", expectedIndex)

       fmt.Println(reader1)
       n, _ = reader1.Read(buf2)
       fmt.Printf("%d bytes were read. (call Read)\n", n)
       fmt.Printf("The reading index in reader: %d\n",
           reader1.Size()-int64(reader1.Len()))
       fmt.Println(reader1)
 */





}
