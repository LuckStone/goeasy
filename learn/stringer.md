# �ַ������ò���
### �ı��ַ����е�ĳЩ�ַ�

str := "hello roc"
bytes := []byte(str)
bytes[1] = 'a'
str = string(bytes) //str == "hallo roc"

### ��ȡ�Ӵ�

substr := str[n:m] //��ȡ������n��m-1���Ӵ�

### �����ַ���

//for�������˷�ʽֻ�ܱ������ASCII
//������ַ������������ľͲ���
for i := 0; i < len(str); i++ {
	//... = str[i]
}
//for-range�������˷�ʽ���Ա���Unicode
//������ַ���������������
for index, char := range str {
	fmt.Printf("%d %c\n",index,char)
}

### �����ַ�������

//�ַ������ַ�ȫΪASCII�е��ַ�
len(str)
//�ַ����к���ASCII��Unicode�ַ�
utf8.RuneCountInString(str)

### �����ַ���

### �ٶ�����д����

var buf bytes.Buffer
buf.WriteString("hello ")
buf.WriteString("roc")
fmt.Println(buf.String()) //hello roc

### �����������ַ�ʽ��

str := strings.Join([]string{"hello"," world"},"")
fmt.Println(str)

str := "hello"
str += "roc"