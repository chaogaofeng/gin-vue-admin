
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="账单编号:" prop="no">
          <el-input v-model="formData.no" :clearable="true"  placeholder="请输入账单编号" />
       </el-form-item>
        <el-form-item label="订单编号:" prop="orderNo">
          <el-input v-model="formData.orderNo" :clearable="true"  placeholder="请输入订单编号" />
       </el-form-item>
        <el-form-item label="账单计划:" prop="plan_id">
        <el-select  v-model="formData.plan_id" placeholder="请选择账单计划" style="width:100%" :clearable="true" >
          <el-option v-for="(item,key) in dataSource.plan_id" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="账单计划数量:" prop="count">
          <el-input v-model.number="formData.count" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="交易哈希:" prop="hash">
          <el-input v-model="formData.hash" :clearable="true"  placeholder="请输入交易哈希" />
       </el-form-item>
        <el-form-item label="开始时间:" prop="start_at">
          <el-date-picker v-model="formData.start_at" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="结束时间:" prop="end_at">
          <el-date-picker v-model="formData.end_at" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="账单状态:" prop="status">
           <el-select v-model="formData.status" placeholder="请选择账单状态" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in bill_statusOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="账单金额:" prop="amount">
          <el-input-number v-model="formData.amount" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="完成时间:" prop="complete_at">
          <el-date-picker v-model="formData.complete_at" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
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
    getBillDataSource,
  createBill,
  updateBill,
  findBill
} from '@/api/swap/bill'

defineOptions({
    name: 'BillForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
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
// 验证规则
const rule = reactive({
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getBillDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findBill({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    bill_statusOptions.value = await getDictFunc('bill_status')
}

init()
// 保存按钮
const save = async() => {
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
