package main

import (
	"database/sql"
	"fmt"
	"strings"
)

import (
	_ "github.com/mattn/go-adodb"
)

type Mssql struct {
	*sql.DB
	dataSource string
	database   string
	windows    bool
	sa         SA
}

type SA struct {
	user   string
	passwd string
}

func (m *Mssql) Open() (err error) {
	var conf []string
	conf = append(conf, "Provider=SQLOLEDB")
	conf = append(conf, "Data Source="+m.dataSource)
	if m.windows {
		// Integrated Security=SSPI �����ʾ�Ե�ǰWINDOWSϵͳ�û���ȥ��¼SQL SERVER������(��Ҫ�ڰ�װsqlserverʱ������)��
		// ���SQL SERVER��������֧�����ַ�ʽ��¼ʱ���ͻ����
		conf = append(conf, "integrated security=SSPI")
	}
	conf = append(conf, "Initial Catalog="+m.database)
	conf = append(conf, "user id="+m.sa.user)
	conf = append(conf, "password="+m.sa.passwd)

	m.DB, err = sql.Open("adodb", strings.Join(conf, ";"))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	db := Mssql{
		dataSource: "CODY\\SQLEXPRESS",
		database:   "test",
		// windwos: true Ϊwindows�����֤��false ��������sa�˺ź�����
		windows: true,
		sa: SA{
			user:   "sa",
			passwd: "123456",
		},
	}
	// �������ݿ�
	err := db.Open()
	if err != nil {
		fmt.Println("sql open:", err)
		return
	}
	defer db.Close()

	// ִ��SQL���
	rows, err := db.Query("select * from info")
	if err != nil {
		fmt.Println("query: ", err)
		return
	}
	for rows.Next() {
		var name string
		var number int
		rows.Scan(&name, &number)
		fmt.Printf("Name: %s \t Number: %d\n", name, number)
	}
}