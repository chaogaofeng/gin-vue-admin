
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
       <el-form-item :label="t('upay.APP.AppId')"  prop="appId" >
          <el-input v-model="formData.appId" :clearable="true"  :placeholder="t('upay.APP.AppId')" />
       </el-form-item>
       <el-form-item :label="t('upay.APP.AppSecret')"  prop="appSecret" >
          <el-input v-model="formData.appSecret" :clearable="true"  :placeholder="t('upay.APP.AppSecret')" />
       </el-form-item>
       <el-form-item :label="t('upay.APP.AppName')"  prop="appName" >
          <el-input v-model="formData.appName" :clearable="true"  :placeholder="t('upay.APP.AppName')" />
       </el-form-item>
       <el-form-item :label="t('upay.APP.Status')"  prop="status" >
           <el-select v-model="formData.status" :placeholder="t('upay.APP.Status')" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in AppStatusOptions" :key="key" :label="item.label" :value="item.value" />
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
    getAPPDataSource,
  createAPP,
  updateAPP,
  findAPP
} from '@/api/upay/app'

defineOptions({
    name: 'APPForm'
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
const AppStatusOptions = ref([])
const formData = ref({
            appId: '',
            appSecret: '',
            appName: '',
            status: '',
        })
// 验证规则
const rule = reactive({
               appId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               appSecret : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               userID : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getAPPDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findAPP({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    AppStatusOptions.value = await getDictFunc('AppStatus')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createAPP(formData.value)
               break
             case 'update':
               res = await updateAPP(formData.value)
               break
             default:
               res = await createAPP(formData.value)
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
