
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
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
    getWalletAddressDataSource,
  createWalletAddress,
  updateWalletAddress,
  findWalletAddress
} from '@/api/upay/wallet_address'

defineOptions({
    name: 'WalletAddressForm'
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
const AddressStatusOptions = ref([])
const formData = ref({
            address: '',
            name: '',
            chainType: '',
            status: '',
        })
// 验证规则
const rule = reactive({
               address : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               chainType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
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
    const res = await getWalletAddressDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findWalletAddress({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    ChainTypeOptions.value = await getDictFunc('ChainType')
    AddressStatusOptions.value = await getDictFunc('AddressStatus')
}

init()
// 保存按钮
const save = async() => {
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
