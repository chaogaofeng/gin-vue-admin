
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
      
        <el-form-item :label="t('upay.WalletAddress.Address')" prop="address">
         <el-input v-model="searchInfo.address" :placeholder="t('general.searchCriteria')" />

        </el-form-item>
        <el-form-item :label="t('upay.WalletAddress.Name')" prop="name">
         <el-input v-model="searchInfo.name" :placeholder="t('general.searchCriteria')" />

        </el-form-item>
           <el-form-item :label="t('upay.WalletAddress.ChainType')" prop="chainType">
            <el-select v-model="searchInfo.chainType" clearable :placeholder="t('general.pleaseSelect')" @clear="()=>{searchInfo.chainType=undefined}">
              <el-option v-for="(item,key) in ChainTypeOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
           <el-form-item :label="t('upay.WalletAddress.Status')" prop="status">
            <el-select v-model="searchInfo.status" clearable :placeholder="t('general.pleaseSelect')" @clear="()=>{searchInfo.status=undefined}">
              <el-option v-for="(item,key) in AddressStatusOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item :label="t('upay.WalletAddress.UserID')" prop="userID">
            
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
            <ExportTemplate  template-id="upay_WalletAddress" />
            <ExportExcel  template-id="upay_WalletAddress" />
            <ImportExcel  template-id="upay_WalletAddress" @on-success="getTableData" />
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" :label="t('general.createdAt')" prop="createdAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
        <el-table-column align="left" :label="t('upay.WalletAddress.Address')" prop="address" width="120" />
        <el-table-column align="left" :label="t('upay.WalletAddress.Name')" prop="name" width="120" />
        <el-table-column align="left" :label="t('upay.WalletAddress.ChainType')" prop="chainType" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.chainType,ChainTypeOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" :label="t('upay.WalletAddress.Status')" prop="status" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.status,AddressStatusOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" :label="t('upay.WalletAddress.UserID')" prop="userID" width="120">
          <template #default="scope">
                
                    <span>{{ filterDataSource(dataSource.userID,scope.row.userID) }}</span>
                
         </template>
         </el-table-column>
        <el-table-column align="left" :label="t('general.operations')" fixed="right" min-width="240">
          <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>{{t('general.desc')}}</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateWalletAddressFunc(scope.row)">{{t('general.change')}}</el-button>
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
            <el-form-item :label="t('upay.WalletAddress.Address')"  prop="address" >
              <el-input v-model="formData.address" :clearable="true"  :placeholder="t('upay.WalletAddress.Address')" />
            </el-form-item>
            <el-form-item :label="t('upay.WalletAddress.Name')"  prop="name" >
              <el-input v-model="formData.name" :clearable="true"  :placeholder="t('upay.WalletAddress.Name')" />
            </el-form-item>
            <el-form-item :label="t('upay.WalletAddress.ChainType')"  prop="chainType" >
              <el-select v-model="formData.chainType" :placeholder="t('upay.WalletAddress.ChainType')" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in ChainTypeOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item :label="t('upay.WalletAddress.Status')"  prop="status" >
              <el-select v-model="formData.status" :placeholder="t('upay.WalletAddress.Status')" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in AddressStatusOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow">
            <el-descriptions :column="1" border>
                    <el-descriptions-item :label="t('upay.WalletAddress.Address')">
                        {{ detailFrom.address }}
                    </el-descriptions-item>
                    <el-descriptions-item :label="t('upay.WalletAddress.Name')">
                        {{ detailFrom.name }}
                    </el-descriptions-item>
                    <el-descriptions-item :label="t('upay.WalletAddress.ChainType')">
                        {{ detailFrom.chainType }}
                    </el-descriptions-item>
                    <el-descriptions-item :label="t('upay.WalletAddress.Status')">
                        {{ detailFrom.status }}
                    </el-descriptions-item>
                    <el-descriptions-item :label="t('upay.WalletAddress.UserID')">
                        {{ detailFrom.userID }}
                    </el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
    getWalletAddressDataSource,
  createWalletAddress,
  deleteWalletAddress,
  deleteWalletAddressByIds,
  updateWalletAddress,
  findWalletAddress,
  getWalletAddressList
} from '@/api/upay/wallet_address'

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
    name: 'WalletAddress'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const ChainTypeOptions = ref([])
const AddressStatusOptions = ref([])
const formData = ref({
            address: '',
            name: '',
            chainType: '',
            status: '',
        })
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getWalletAddressDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()



// 验证规则
const rule = reactive({
               address : [{
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
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

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
  const table = await getWalletAddressList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    AddressStatusOptions.value = await getDictFunc('AddressStatus')
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
            deleteWalletAddressFunc(row)
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
      const res = await deleteWalletAddressByIds({ IDs })
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
const updateWalletAddressFunc = async(row) => {
    const res = await findWalletAddress({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteWalletAddressFunc = async (row) => {
    const res = await deleteWalletAddress({ ID: row.ID })
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
        address: '',
        name: '',
        chainType: '',
        status: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createWalletAddress(formData.value)
                  break
                case 'update':
                  res = await updateWalletAddress(formData.value)
                  break
                default:
                  res = await createWalletAddress(formData.value)
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
  const res = await findWalletAddress({ ID: row.ID })
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
