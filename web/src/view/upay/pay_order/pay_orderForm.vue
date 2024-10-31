
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
       <el-form-item :label="t('upay.PayOrder.Status')"  prop="status" >
           <el-select v-model="formData.status" :placeholder="t('upay.PayOrder.Status')" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in PayStatusOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
       <el-form-item :label="t('upay.PayOrder.Amount')"  prop="crypto" >
          <el-input-number v-model="formData.crypto" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
       <el-form-item :label="t('upay.PayOrder.ActualAmount')"  prop="actualCrypto" >
          <el-input-number v-model="formData.actualCrypto" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
       <el-form-item :label="t('upay.PayOrder.Fee')"  prop="poundage" >
          <el-input-number v-model="formData.poundage" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
       <el-form-item :label="t('upay.PayOrder.ActualFee')"  prop="actualPoundage" >
          <el-input-number v-model="formData.actualPoundage" :precision="2" :clearable="true"></el-input-number>
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
          <el-date-picker v-model="formData.completedAt" type="date" :placeholder="t('general.selectDate')" :clearable="true"></el-date-picker>
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
        <el-form-item>
          <el-button type="primary" @click="save">{{ t('general.save') }}</el-button>
          <el-button type="primary" @click="back">{{ t('general.back') }}</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
    getPayOrderDataSource,
  createPayOrder,
  updatePayOrder,
  findPayOrder
} from '@/api/upay/pay_order'

defineOptions({
    name: 'PayOrderForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()


const route = useRoute()
const router = useRouter()

const type = ref('')
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
// 验证规则
const rule = reactive({
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               crypto : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               chainType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               appId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               receiver : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               orderNo : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               merchantOrderNo : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               userID : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getPayOrderDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findPayOrder({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    ChainTypeOptions.value = await getDictFunc('ChainType')
    RiskLevelOptions.value = await getDictFunc('RiskLevel')
    PayStatusOptions.value = await getDictFunc('PayStatus')
}

init()
// 保存按钮
const save = async() => {
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
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
