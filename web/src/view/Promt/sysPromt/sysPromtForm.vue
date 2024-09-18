
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="提示词:" prop="promtData">
          <SelectFile v-model="formData.promtData" />
       </el-form-item>
        <el-form-item label="管理用户ID:" prop="systemUserId">
          <el-input v-model.number="formData.systemUserId" :clearable="true" placeholder="请输入" />
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
  createPromt,
  updatePromt,
  findPromt
} from '@/api/Promt/sysPromt'

defineOptions({
    name: 'PromtForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 文件选择组件
import SelectFile from '@/components/selectFile/selectFile.vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            promtData: [],
            systemUserId: undefined,
        })
// 验证规则
const rule = reactive({
               promtData : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               systemUserId : [{
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
      const res = await findPromt({ ID: route.query.id })
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
               res = await createPromt(formData.value)
               break
             case 'update':
               res = await updatePromt(formData.value)
               break
             default:
               res = await createPromt(formData.value)
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
