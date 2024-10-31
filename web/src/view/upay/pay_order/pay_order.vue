
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item :label="t('general.createDate')" prop="createdAt">
      <template #label>
        <span>
          {{t('general.createDate')}}
          <el-tooltip :content="t('general.searchDesc')">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" :placeholder="t('general.startData')" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" :placeholder="t('general.endData')" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
      
           <el-form-item :label="t('upay.PayOrder.Status')" prop="status">
            <el-select v-model="searchInfo.status" clearable :placeholder="t('general.pleaseSelect')" @clear="()=>{searchInfo.status=undefined}">
              <el-option v-for="(item,key) in PayStatusOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item :label="t('upay.PayOrder.Amount')" prop="crypto">
            
             <el-input v-model.number="searchInfo.crypto" :placeholder="t('general.searchCriteria')" />

        </el-form-item>
        <el-form-item :label="t('upay.PayOrder.ActualAmount')" prop="actualCrypto">
            
             <el-input v-model.number="searchInfo.actualCrypto" :placeholder="t('general.searchCriteria')" />

        </el-form-item>
           <el-form-item :label="t('upay.PayOrder.ChainType')" prop="chainType">
            <el-select v-model="searchInfo.chainType" clearable :placeholder="t('general.pleaseSelect')" @clear="()=>{searchInfo.chainType=undefined}">
              <el-option v-for="(item,key) in ChainTypeOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item :label="t('upay.PayOrder.AppID')" prop="appId">
         <el-input v-model="searchInfo.appId" :placeholder="t('general.searchCriteria')" />

        </el-form-item>
        <el-form-item :label="t('upay.PayOrder.Receiver')" prop="receiver">
         <el-input v-model="searchInfo.receiver" :placeholder="t('general.searchCriteria')" />

        </el-form-item>
        <el-form-item :label="t('upay.PayOrder.Hash')" prop="hash">
         <el-input v-model="searchInfo.hash" :placeholder="t('general.searchCriteria')" />

        </el-form-item>
        <el-form-item :label="t('upay.PayOrder.OrderNo')" prop="orderNo">
         <el-input v-model="searchInfo.orderNo" :placeholder="t('general.searchCriteria')" />

        </el-form-item>
        <el-form-item :label="t('upay.PayOrder.MerchantOrderNo')" prop="merchantOrderNo">
         <el-input v-model="searchInfo.merchantOrderNo" :placeholder="t('general.searchCriteria')" />

        </el-form-item>
        <el-form-item :label="t('upay.PayOrder.CompletedAt')" prop="completedAt">
            
            <template #label>
            <span>
              {{ t('upay.PayOrder.CompletedAt') }}
              <el-tooltip :content="t('general.searchDesc')">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
            <el-date-picker v-model="searchInfo.startCompletedAt" type="datetime" :placeholder="t('general.startData')" :disabled-date="time=> searchInfo.endCompletedAt ? time.getTime() > searchInfo.endCompletedAt.getTime() : false"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.endCompletedAt" type="datetime" :placeholder="t('general.endData')" :disabled-date="time=> searchInfo.startCompletedAt ? time.getTime() < searchInfo.startCompletedAt.getTime() : false"></el-date-picker>

        </el-form-item>
           <el-form-item :label="t('upay.PayOrder.RiskLevel')" prop="riskLevel">
            <el-select v-model="searchInfo.riskLevel" clearable :placeholder="t('general.pleaseSelect')" @clear="()=>{searchInfo.riskLevel=undefined}">
              <el-option v-for="(item,key) in RiskLevelOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item :label="t('upay.PayOrder.UserID')" prop="userID">
            
             <el-input v-model.number="searchInfo.userID" :placeholder="t('general.searchCriteria')" />

        </el-form-item>

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">{{t('general.search')}}</el-button>
          <el-button icon="refresh" @click="onReset">{{t('general.reset')}}</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">{{t('general.expand')}}</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>{{t('general.collapse')}}</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog">{{t('general.add')}}</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">{{t('general.delete')}}</el-button>
            <ExportTemplate  template-id="upay_PayOrder" />
            <ExportExcel  template-id="upay_PayOrder" />
            <ImportExcel  template-id="upay_PayOrder" @on-success="getTableData" />
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
        
        <el-table-column align="left" :label="t('general.createdAt')" prop="createdAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
        <el-table-column align="left" :label="t('upay.PayOrder.Status')" prop="status" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.status,PayStatusOptions) }}
            </template>
        </el-table-column>
        <el-table-column sortable align="left" :label="t('upay.PayOrder.Amount')" prop="crypto" width="120" />
        <el-table-column sortable align="left" :label="t('upay.PayOrder.ActualAmount')" prop="actualCrypto" width="120" />
        <el-table-column align="left" :label="t('upay.PayOrder.Fee')" prop="poundage" width="120" />
        <el-table-column align="left" :label="t('upay.PayOrder.ActualFee')" prop="actualPoundage" width="120" />
        <el-table-column align="left" :label="t('upay.PayOrder.ChainType')" prop="chainType" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.chainType,ChainTypeOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" :label="t('upay.PayOrder.AppID')" prop="appId" width="120" />
        <el-table-column align="left" :label="t('upay.PayOrder.Receiver')" prop="receiver" width="120" />
        <el-table-column align="left" :label="t('upay.PayOrder.Hash')" prop="hash" width="120" />
        <el-table-column align="left" :label="t('upay.PayOrder.OrderNo')" prop="orderNo" width="120" />
        <el-table-column align="left" :label="t('upay.PayOrder.MerchantOrderNo')" prop="merchantOrderNo" width="120" />
        <el-table-column align="left" :label="t('upay.PayOrder.Attach')" prop="attach" width="120" />
         <el-table-column align="left" :label="t('upay.PayOrder.CompletedAt')" prop="completedAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.completedAt) }}</template>
         </el-table-column>
        <el-table-column align="left" :label="t('upay.PayOrder.RiskLevel')" prop="riskLevel" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.riskLevel,RiskLevelOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" :label="t('upay.PayOrder.UserID')" prop="userID" width="120">
          <template #default="scope">
                
                    <span>{{ filterDataSource(dataSource.userID,scope.row.userID) }}</span>
                
         </template>
         </el-table-column>
        <el-table-column align="left" :label="t('general.operations')" fixed="right" min-width="240">
          <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>{{t('general.desc')}}</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updatePayOrderFunc(scope.row)">{{t('general.change')}}</el-button>
            <el-button  type="primary" link icon="delete" @click="deleteRow(scope.row)">{{t('general.delete')}}</el-button>
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
                <span class="text-lg">{{type==='create'?t('general.add'):t('general.edit')}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">{{t('general.confirm')}}</el-button>
                  <el-button @click="closeDialog">{{t('general.close')}}</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item :label="t('upay.PayOrder.Status')"  prop="status" >
              <el-select v-model="formData.status" :placeholder="t('upay.PayOrder.Status')" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in PayStatusOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item :label="t('upay.PayOrder.Amount')"  prop="crypto" >
              <el-input-number v-model="formData.crypto"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item :label="t('upay.PayOrder.ActualAmount')"  prop="actualCrypto" >
              <el-input-number v-model="formData.actualCrypto"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item :label="t('upay.PayOrder.Fee')"  prop="poundage" >
              <el-input-number v-model="formData.poundage"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item :label="t('upay.PayOrder.ActualFee')"  prop="actualPoundage" >
              <el-input-number v-model="formData.actualPoundage"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item :label="t('upay.PayOrder.ChainType')"  prop="chainType" >
              <el-select v-model="formData.chainType" :placeholder="t('upay.PayOrder.ChainType')" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in ChainTypeOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item :label="t('upay.PayOrder.AppID')"  prop="appId" >
              <el-input v-model="formData.appId" :clearable="true"  :placeholder="t('upay.PayOrder.AppID')" />
            </el-form-item>
            <el-form-item :label="t('upay.PayOrder.Receiver')"  prop="receiver" >
              <el-input v-model="formData.receiver" :clearable="true"  :placeholder="t('upay.PayOrder.Receiver')" />
            </el-form-item>
            <el-form-item :label="t('upay.PayOrder.Hash')"  prop="hash" >
              <el-input v-model="formData.hash" :clearable="true"  :placeholder="t('upay.PayOrder.Hash')" />
            </el-form-item>
            <el-form-item :label="t('upay.PayOrder.OrderNo')"  prop="orderNo" >
              <el-input v-model="formData.orderNo" :clearable="true"  :placeholder="t('upay.PayOrder.OrderNo')" />
            </el-form-item>
            <el-form-item :label="t('upay.PayOrder.MerchantOrderNo')"  prop="merchantOrderNo" >
              <el-input v-model="formData.merchantOrderNo" :clearable="true"  :placeholder="t('upay.PayOrder.MerchantOrderNo')" />
            </el-form-item>
            <el-form-item :label="t('upay.PayOrder.Attach')"  prop="attach" >
              <el-input v-model="formData.attach" :clearable="true"  :placeholder="t('upay.PayOrder.Attach')" />
            </el-form-item>
            <el-form-item :label="t('upay.PayOrder.CompletedAt')"  prop="completedAt" >
              <el-date-picker v-model="formData.completedAt" type="date" style="width:100%" :placeholder="t('general.selectData')" :clearable="true"  />
            </el-form-item>
            <el-form-item :label="t('upay.PayOrder.RiskLevel')"  prop="riskLevel" >
              <el-select v-model="formData.riskLevel" :placeholder="t('upay.PayOrder.RiskLevel')" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in RiskLevelOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item :label="t('upay.PayOrder.UserID')"  prop="userID" >
            <el-select  v-model="formData.userID" :placeholder="t('upay.PayOrder.UserID')" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in dataSource.userID" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow">
            <el-descriptions :column="1" border>
                    <el-descriptions-item :label="t('upay.PayOrder.Status')">
                        {{ detailFrom.status }}
                    </el-descriptions-item>
                    <el-descriptions-item :label="t('upay.PayOrder.Amount')">
                        {{ detailFrom.crypto }}
                    </el-descriptions-item>
                    <el-descriptions-item :label="t('upay.PayOrder.ActualAmount')">
                        {{ detailFrom.actualCrypto }}
                    </el-descriptions-item>
                    <el-descriptions-item :label="t('upay.PayOrder.Fee')">
                        {{ detailFrom.poundage }}
                    </el-descriptions-item>
                    <el-descriptions-item :label="t('upay.PayOrder.ActualFee')">
                        {{ detailFrom.actualPoundage }}
                    </el-descriptions-item>
                    <el-descriptions-item :label="t('upay.PayOrder.ChainType')">
                        {{ detailFrom.chainType }}
                    </el-descriptions-item>
                    <el-descriptions-item :label="t('upay.PayOrder.AppID')">
                        {{ detailFrom.appId }}
                    </el-descriptions-item>
                    <el-descriptions-item :label="t('upay.PayOrder.Receiver')">
                        {{ detailFrom.receiver }}
                    </el-descriptions-item>
                    <el-descriptions-item :label="t('upay.PayOrder.Hash')">
                        {{ detailFrom.hash }}
                    </el-descriptions-item>
                    <el-descriptions-item :label="t('upay.PayOrder.OrderNo')">
                        {{ detailFrom.orderNo }}
                    </el-descriptions-item>
                    <el-descriptions-item :label="t('upay.PayOrder.MerchantOrderNo')">
                        {{ detailFrom.merchantOrderNo }}
                    </el-descriptions-item>
                    <el-descriptions-item :label="t('upay.PayOrder.Attach')">
                        {{ detailFrom.attach }}
                    </el-descriptions-item>
                    <el-descriptions-item :label="t('upay.PayOrder.CompletedAt')">
                        {{ detailFrom.completedAt }}
                    </el-descriptions-item>
                    <el-descriptions-item :label="t('upay.PayOrder.RiskLevel')">
                        {{ detailFrom.riskLevel }}
                    </el-descriptions-item>
                    <el-descriptions-item :label="t('upay.PayOrder.UserID')">
                        {{ detailFrom.userID }}
                    </el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
    getPayOrderDataSource,
  createPayOrder,
  deletePayOrder,
  deletePayOrderByIds,
  updatePayOrder,
  findPayOrder,
  getPayOrderList
} from '@/api/upay/pay_order'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'

defineOptions({
    name: 'PayOrder'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const ChainTypeOptions = ref([])
const RiskLevelOptions = ref([])
const PayStatusOptions = ref([])
const formData = ref({
            status: '',
            crypto: 0,
            actualCrypto: 0,
            poundage: 0,
            actualPoundage: 0,
            chainType: '',
            appId: '',
            receiver: '',
            hash: '',
            orderNo: '',
            merchantOrderNo: '',
            attach: '',
            completedAt: new Date(),
            riskLevel: '',
            userID: undefined,
        })
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getPayOrderDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()



// 验证规则
const rule = reactive({
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: t('general.noOnlySpace'),
                   trigger: ['input', 'blur'],
              }
              ],
               crypto : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               chainType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: t('general.noOnlySpace'),
                   trigger: ['input', 'blur'],
              }
              ],
               appId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: t('general.noOnlySpace'),
                   trigger: ['input', 'blur'],
              }
              ],
               receiver : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: t('general.noOnlySpace'),
                   trigger: ['input', 'blur'],
              }
              ],
               orderNo : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: t('general.noOnlySpace'),
                   trigger: ['input', 'blur'],
              }
              ],
               merchantOrderNo : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: t('general.noOnlySpace'),
                   trigger: ['input', 'blur'],
              }
              ],
               userID : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error(t('general.placeInputEndData')))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error(t('general.placeInputStartData')))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error(t('general.startDataMustBeforeEndData')))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
        completedAt : [{ validator: (rule, value, callback) => {
        if (searchInfo.value.startCompletedAt && !searchInfo.value.endCompletedAt) {
          callback(new Error(t('general.placeInputEndData')))
        } else if (!searchInfo.value.startCompletedAt && searchInfo.value.endCompletedAt) {
          callback(new Error(t('general.placeInputStartData')))
        } else if (searchInfo.value.startCompletedAt && searchInfo.value.endCompletedAt && (searchInfo.value.startCompletedAt.getTime() === searchInfo.value.endCompletedAt.getTime() || searchInfo.value.startCompletedAt.getTime() > searchInfo.value.endCompletedAt.getTime())) {
          callback(new Error(t('general.startDataMustBeforeEndData')))
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
            crypto: 'amount',
            actualCrypto: 'actual_amount',
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
  const table = await getPayOrderList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    ChainTypeOptions.value = await getDictFunc('ChainType')
    RiskLevelOptions.value = await getDictFunc('RiskLevel')
    PayStatusOptions.value = await getDictFunc('PayStatus')
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
  ElMessageBox.confirm(t('general.deleteConfirm'), t('general.hint'), {
    confirmButtonText: t('general.confirm'),
    cancelButtonText: t('general.cancel'),
        type: 'warning'
    }).then(() => {
            deletePayOrderFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm(t('general.deleteConfirm'), t('general.hint'), {
    confirmButtonText: t('general.confirm'),
    cancelButtonText: t('general.cancel'),
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: t('general.selectDataToDelete')
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deletePayOrderByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: t('general.deleteSuccess')
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
const updatePayOrderFunc = async(row) => {
    const res = await findPayOrder({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deletePayOrderFunc = async (row) => {
    const res = await deletePayOrder({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: t('general.deleteSuccess')
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
        status: '',
        crypto: 0,
        actualCrypto: 0,
        poundage: 0,
        actualPoundage: 0,
        chainType: '',
        appId: '',
        receiver: '',
        hash: '',
        orderNo: '',
        merchantOrderNo: '',
        attach: '',
        completedAt: new Date(),
        riskLevel: '',
        userID: undefined,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createPayOrder(formData.value)
                  break
                case 'update':
                  res = await updatePayOrder(formData.value)
                  break
                default:
                  res = await createPayOrder(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: t('general.createUpdateSuccess')
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
  const res = await findPayOrder({ ID: row.ID })
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
