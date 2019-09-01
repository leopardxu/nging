// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo/param"
	
	"time"
)

// TaskLog 任务日志
type TaskLog struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*TaskLog
	namer   func(string) string
	connID  int
	
	Id     	uint64  	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"" json:"id" xml:"id"`
	TaskId 	uint    	`db:"task_id" bson:"task_id" comment:"任务ID" json:"task_id" xml:"task_id"`
	Output 	string  	`db:"output" bson:"output" comment:"任务输出" json:"output" xml:"output"`
	Error  	string  	`db:"error" bson:"error" comment:"错误信息" json:"error" xml:"error"`
	Status 	string  	`db:"status" bson:"status" comment:"状态" json:"status" xml:"status"`
	Elapsed	uint    	`db:"elapsed" bson:"elapsed" comment:"消耗时间(毫秒)" json:"elapsed" xml:"elapsed"`
	Created	uint    	`db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
}

func (this *TaskLog) Trans() *factory.Transaction {
	return this.trans
}

func (this *TaskLog) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *TaskLog) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *TaskLog) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *TaskLog) Objects() []*TaskLog {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *TaskLog) NewObjects() *[]*TaskLog {
	this.objects = []*TaskLog{}
	return &this.objects
}

func (this *TaskLog) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *TaskLog) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *TaskLog) Short_() string {
	return "task_log"
}

func (this *TaskLog) Struct_() string {
	return "TaskLog"
}

func (this *TaskLog) Name_() string {
	if this.namer != nil {
		return WithPrefix(this.namer(this.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(this.Short_())(this))
}

func (this *TaskLog) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *TaskLog) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *TaskLog) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *TaskLog) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *TaskLog) GroupBy(keyField string, inputRows ...[]*TaskLog) map[string][]*TaskLog {
	var rows []*TaskLog
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string][]*TaskLog{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*TaskLog{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (this *TaskLog) KeyBy(keyField string, inputRows ...[]*TaskLog) map[string]*TaskLog {
	var rows []*TaskLog
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]*TaskLog{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (this *TaskLog) AsKV(keyField string, valueField string, inputRows ...[]*TaskLog) map[string]interface{} {
	var rows []*TaskLog
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]interface{}{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (this *TaskLog) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *TaskLog) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Status) == 0 { this.Status = "success" }
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint64(v)
		}
	}
	return
}

func (this *TaskLog) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	
	if len(this.Status) == 0 { this.Status = "success" }
	return this.Setter(mw, args...).SetSend(this).Update()
}

func (this *TaskLog) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *TaskLog) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) error {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *TaskLog) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) error {
	
	if val, ok := kvset["status"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["status"] = "success" } }
	return this.Setter(mw, args...).SetSend(kvset).Update()
}

func (this *TaskLog) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		
	if len(this.Status) == 0 { this.Status = "success" }
	},func(){
		this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Status) == 0 { this.Status = "success" }
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint64(v)
		}
	}
	return 
}

func (this *TaskLog) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

func (this *TaskLog) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *TaskLog) Reset() *TaskLog {
	this.Id = 0
	this.TaskId = 0
	this.Output = ``
	this.Error = ``
	this.Status = ``
	this.Elapsed = 0
	this.Created = 0
	return this
}

func (this *TaskLog) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["TaskId"] = this.TaskId
	r["Output"] = this.Output
	r["Error"] = this.Error
	r["Status"] = this.Status
	r["Elapsed"] = this.Elapsed
	r["Created"] = this.Created
	return r
}

func (this *TaskLog) Set(key interface{}, value ...interface{}) {
	switch k := key.(type) {
		case map[string]interface{}:
			for kk, vv := range k {
				this.Set(kk, vv)
			}
		default:
			var (
				kk string
				vv interface{}
			)
			if k, y := key.(string); y {
				kk = k
			} else {
				kk = fmt.Sprint(key)
			}
			if len(value) > 0 {
				vv = value[0]
			}
			switch kk {
				case "Id": this.Id = param.AsUint64(vv)
				case "TaskId": this.TaskId = param.AsUint(vv)
				case "Output": this.Output = param.AsString(vv)
				case "Error": this.Error = param.AsString(vv)
				case "Status": this.Status = param.AsString(vv)
				case "Elapsed": this.Elapsed = param.AsUint(vv)
				case "Created": this.Created = param.AsUint(vv)
			}
	}
}

func (this *TaskLog) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = this.Id
	r["task_id"] = this.TaskId
	r["output"] = this.Output
	r["error"] = this.Error
	r["status"] = this.Status
	r["elapsed"] = this.Elapsed
	r["created"] = this.Created
	return r
}

func (this *TaskLog) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = this.AsRow()
	}
	return factory.BatchValidate(this.Short_(), kvset)
}

func (this *TaskLog) Validate(field string, value interface{}) error {
	return factory.Validate(this.Short_(), field, value)
}

