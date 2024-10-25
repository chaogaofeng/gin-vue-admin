
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
      
        <el-form-item label="账单编号" prop="no">
         <el-input v-model="searchInfo.no" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="订单编号" prop="orderNo">
         <el-input v-model="searchInfo.orderNo" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="账单计划" prop="plan_id">
            
             <el-input v-model.number="searchInfo.plan_id" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="交易哈希" prop="hash">
         <el-input v-model="searchInfo.hash" placeholder="搜索条件" />

        </el-form-item>
           <el-form-item label="账单状态" prop="status">
            <el-select v-model="searchInfo.status" clearable placeholder="请选择" @clear="()=>{searchInfo.status=undefined}">
              <el-option v-for="(item,key) in bill_statusOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="完成时间" prop="complete_at">
            
            <template #label>
            <span>
              完成时间
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
            <el-date-picker v-model="searchInfo.startCompleteAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCompleteAt ? time.getTime() > searchInfo.endCompleteAt.getTime() : false"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.endCompleteAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCompleteAt ? time.getTime() < searchInfo.startCompleteAt.getTime() : false"></el-date-picker>

        </el-form-item>

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="日期" prop="createdAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
          <el-table-column align="left" label="账单编号" prop="no" width="120" />
          <el-table-column align="left" label="订单编号" prop="orderNo" width="120" />
        <el-table-column sortable align="left" label="账单计划" prop="plan_id" width="120">
          <template #default="scope">
                
                    <span>{{ filterDataSource(dataSource.plan_id,scope.row.plan_id) }}</span>
                
         </template>
         </el-table-column>
          <el-table-column sortable align="left" label="账单计划数量" prop="count" width="120" />
          <el-table-column align="left" label="交易哈希" prop="hash" width="120" />
         <el-table-column align="left" label="开始时间" prop="start_at" width="180">
            <template #default="scope">{{ formatDate(scope.row.start_at) }}</template>
         </el-table-column>
         <el-table-column align="left" label="结束时间" prop="end_at" width="180">
            <template #default="scope">{{ formatDate(scope.row.end_at) }}</template>
         </el-table-column>
        <el-table-column sortable align="left" label="账单状态" prop="status" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.status,bill_statusOptions) }}
            </template>
        </el-table-column>
          <el-table-column sortable align="left" label="账单金额" prop="amount" width="120" />
         <el-table-column sortable align="left" label="完成时间" prop="complete_at" width="180">
            <template #default="scope">{{ formatDate(scope.row.complete_at) }}</template>
         </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看详情</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateBillFunc(scope.row)">变更</el-button>
            <el-button  type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'添加':'修改'}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="账单编号:"  prop="no" >
              <el-input v-model="formData.no" :clearable="true"  placeholder="请输入账单编号" />
            </el-form-item>
            <el-form-item label="订单编号:"  prop="orderNo" >
              <el-input v-model="formData.orderNo" :clearable="true"  placeholder="请输入订单编号" />
            </el-form-item>
            <el-form-item label="账单计划:"  prop="plan_id" >
            <el-select  v-model="formData.plan_id" placeholder="请选择账单计划" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in dataSource.plan_id" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
            <el-form-item label="账单计划数量:"  prop="count" >
              <el-input v-model.number="formData.count" :clearable="true" placeholder="请输入账单计划数量" />
            </el-form-item>
            <el-form-item label="交易哈希:"  prop="hash" >
              <el-input v-model="formData.hash" :clearable="true"  placeholder="请输入交易哈希" />
            </el-form-item>
            <el-form-item label="开始时间:"  prop="start_at" >
              <el-date-picker v-model="formData.start_at" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="结束时间:"  prop="end_at" >
              <el-date-picker v-model="formData.end_at" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="账单状态:"  prop="status" >
              <el-select v-model="formData.status" placeholder="请选择账单状态" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in bill_statusOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="账单金额:"  prop="amount" >
              <el-input-number v-model="formData.amount"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="完成时间:"  prop="complete_at" >
              <el-date-picker v-model="formData.complete_at" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="账单编号">
                        {{ detailFrom.no }}
                    </el-descriptions-item>
                    <el-descriptions-item label="订单编号">
                        {{ detailFrom.orderNo }}
                    </el-descriptions-item>
                    <el-descriptions-item label="账单计划">
                        {{ detailFrom.plan_id }}
                    </el-descriptions-item>
                    <el-descriptions-item label="账单计划数量">
                        {{ detailFrom.count }}
                    </el-descriptions-item>
                    <el-descriptions-item label="交易哈希">
                        {{ detailFrom.hash }}
                    </el-descriptions-item>
                    <el-descriptions-item label="开始时间">
                        {{ detailFrom.start_at }}
                    </el-descriptions-item>
                    <el-descriptions-item label="结束时间">
                        {{ detailFrom.end_at }}
                    </el-descriptions-item>
                    <el-descriptions-item label="账单状态">
                        {{ detailFrom.status }}
                    </el-descriptions-item>
                    <el-descriptions-item label="账单金额">
                        {{ detailFrom.amount }}
                    </el-descriptions-item>
                    <el-descriptions-item label="完成时间">
                        {{ detailFrom.complete_at }}
                    </el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
    getBillDataSource,
  createBill,
  deleteBill,
  deleteBillByIds,
  updateBill,
  findBill,
  getBillList
} from '@/api/swap/bill'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'




defineOptions({
    name: 'Bill'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const bill_statusOptions = ref([])
const formData = ref({
            no: '',
            orderNo: '',
            plan_id: undefined,
            count: undefined,
            hash: '',
            start_at: new Date(),
            end_at: new Date(),
            status: '',
            amount: 0,
            complete_at: new Date(),
        })
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getBillDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()



// 验证规则
const rule = reactive({
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
        complete_at : [{ validator: (rule, value, callback) => {
        if (searchInfo.value.startCompleteAt && !searchInfo.value.endCompleteAt) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startCompleteAt && searchInfo.value.endCompleteAt) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startCompleteAt && searchInfo.value.endCompleteAt && (searchInfo.value.startCompleteAt.getTime() === searchInfo.value.endCompleteAt.getTime() || searchInfo.value.startCompleteAt.getTime() > searchInfo.value.endCompleteAt.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change' }],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
            plan_id: 'plan_id',
            count: 'count',
            status: 'status',
            amount: 'amount',
            complete_at: 'complete_at',
  }

  let sort = sortMap[prop]
  if(!sort){
   sort = prop.replace(/[A-Z]/g, match => `_${match.toLowerCase()}`)
  }

  searchInfo.value.sort = sort
  searchInfo.value.order = order
  getTableData()
}

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getBillList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    bill_statusOptions.value = await getDictFunc('bill_status')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteBillFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteBillByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateBillFunc = async(row) => {
    const res = await findBill({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteBillFunc = async (row) => {
    const res = await deleteBill({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        no: '',
        orderNo: '',
        plan_id: undefined,
        count: undefined,
        hash: '',
        start_at: new Date(),
        end_at: new Date(),
        status: '',
        amount: 0,
        complete_at: new Date(),
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createBill(formData.value)
                  break
                case 'update':
                  res = await updateBill(formData.value)
                  break
                default:
                  res = await createBill(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}


const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findBill({ ID: row.ID })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}


</script>

<style>

</style>