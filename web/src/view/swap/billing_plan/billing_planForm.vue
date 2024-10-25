
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入名称" />
       </el-form-item>
        <el-form-item label="价格:" prop="price">
          <el-input-number v-model="formData.price" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="数量:" prop="duration">
          <el-input v-model.number="formData.duration" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="单位:" prop="durationUnit">
        <el-select v-model="formData.durationUnit" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in ['D','W','M','Q','Y']" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="日价格:" prop="dailyPrice">
          <el-input-number v-model="formData.dailyPrice" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createBillingPlan,
  updateBillingPlan,
  findBillingPlan
} from '@/api/swap/billing_plan'

defineOptions({
    name: 'BillingPlanForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            name: '',
            price: 0,
            duration: undefined,
            dailyPrice: 0,
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               price : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               duration : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               durationUnit : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findBillingPlan({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createBillingPlan(formData.value)
               break
             case 'update':
               res = await updateBillingPlan(formData.value)
               break
             default:
               res = await createBillingPlan(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
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
