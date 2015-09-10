package main

import (
	"fmt"
	"github.com/ruiaylin/sqlparser"
	//"github.com/ruiaylin/sqlparser/dependency/sqltypes"
)

func main() {
	TestParse()
}

func TestParse() {
	fmt.Println(">>>>> TestParse() ")
	//sql := "select a from (select * from table1 where table1.a = 'tom') as t1, table2, table3 as t3, table4 left join table5 where t1.k = '1'"
	sql1 := " select * from table1 where table1.a = 'tom' "
	xd, err := sqlparser.Parse(sql1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(">>>>>: ", sqlparser.String(xd))
	fmt.Println(">>>>>: type  ")

	switch node := xd.(type) {
	case *sqlparser.Select:
		fmt.Printf("select %v from %v where 1 != 1", node.SelectExprs, node.From)
	default:
		fmt.Printf("test\n")
	}
	fmt.Println("-------------------------------------------------------")
	sql2 := `create table t1 (
	LastName varchar(255),
	FirstName varchar(255),
	ID int primary key
	)`
	tree, err := sqlparser.Parse(sql2)
	primary_key, err := sqlparser.GetPrimaryKey(tree)
	fmt.Println(" primary key ", primary_key.ColName)

	sq := `INSERT INTO statsinfos(
        daytime ,ipaddr,port ,hostname , load1m ,memt , memf , 
        sys ,wio ,user ,sess ,act , ins  ,upd  ,dele ,sel, 
        hito,otrx,rowi,rowu,rowd,rows,data,redo
        ) 
    VALUES ('2015-08-27 19:42:40.437768','192.168.7.128','3306',  'BJ-YDZC-7-128', '0.01', 258443.03, 239513.75,
            0,0,0,  2 ,2, 0, 0,0,0 , 
            100,0,   0,0,0,0  , 0 , 0  )`
	sq1 := `select id,item_id,city_id,coupon_channel,parallel_import_channel,installment_channel,reassure_channel,brand_id,brand_name,series_id, series_name,level_id,level_name,product_price,price_range,price_range_id,min_price,max_price,brand_order,series_order, level_order,last_update_time,status,factory_id,factory_name,factory_order,pre_time,start_time,end_time from item_city_relation_0 r where 1=1 and pre_time<now() and end_time>now() and brand_id = 81 group by brand_id, factory_id, series_id, level_id, price_range order by brand_order asc, level_order asc, series_order desc, factory_order asc  `
	t1, err := sqlparser.Parse(sq)
	typeSwitch(t1)
	t2, err := sqlparser.Parse(sq1)
	typeSwitch(t2)
}

func typeSwitch(node sqlparser.SQLNode) {
	switch node := node.(type) {
	case *sqlparser.Select:
		fmt.Printf("select %v from %v where 1 != 1", sqlparser.String(node.SelectExprs), sqlparser.String(node.From))
		fmt.Println("node.Comments : ", sqlparser.String(node.Comments), "\n")
		fmt.Println("node.Distinct : ", node.Distinct, "\n")
		fmt.Println("node.SelectExpr : ", sqlparser.String(node.SelectExprs), "\n")
		fmt.Println("node.From : ", sqlparser.String(node.From), "\n")
		fmt.Println("node.Wher : ", sqlparser.String(node.Where), "\n")
		fmt.Println("node.GroupB : ", sqlparser.String(node.GroupBy), "\n")
		fmt.Println("node.Having : ", sqlparser.String(node.Having), "\n")
		fmt.Println("node.OrderB : ", sqlparser.String(node.OrderBy), "\n")
		fmt.Println("node.Limit : ", sqlparser.String(node.Limit), "\n")
		fmt.Println("node.Lock : ", node.Lock, "\n")
	case *sqlparser.Insert:
		fmt.Println("node.Comments : ", sqlparser.String(node.Comments), "\n")
		fmt.Println("node.Table : ", sqlparser.String(node.Table), "\n")
		fmt.Println("node.Columns : ", sqlparser.String(node.Columns), "\n")
		fmt.Println("node.Rows : ", sqlparser.String(node.Rows), "\n")
		fmt.Println("node.OnDup : ", sqlparser.String(node.OnDup), "\n")
	case *sqlparser.Update:
		fmt.Println("node.Comments : ", sqlparser.String(node.Comments), "\n ")
		fmt.Println("node.Table : ", sqlparser.String(node.Table), "\n ")
		fmt.Println("node.Exprs : ", sqlparser.String(node.Exprs), "\n ")
		fmt.Println("node.Where : ", sqlparser.String(node.Where), "\n ")
		fmt.Println("node.OrderBy : ", sqlparser.String(node.OrderBy), "\n ")
		fmt.Println("node.Limit : ", sqlparser.String(node.Limit), "\n ")
	case *sqlparser.Delete:
		fmt.Println("node.Comments : ", sqlparser.String(node.Comments), "\n")
		fmt.Println("node.Table : ", sqlparser.String(node.Table), "\n")
		fmt.Println("node.Where : ", sqlparser.String(node.Where), "\n")
		fmt.Println("node.OrderBy : ", sqlparser.String(node.OrderBy), "\n")
		fmt.Println("node.Limit : ", sqlparser.String(node.Limit), "\n")
	default:
		fmt.Printf("test\n")
	}

}
