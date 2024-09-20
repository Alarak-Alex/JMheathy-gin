<template>
    <div>
        <el-upload :action="`${getBaseUrl()}/fileUploadAndDownload/upload?noSave=1`" :on-change="onChange"
            :show-file-list="true" :limit="limit" :accept="accept" class="upload-btn">
            <el-button type="primary">
                上传文件
            </el-button>
        </el-upload>
    </div>
</template>

<script setup>
import { ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { getBaseUrl } from '@/utils/format'
import * as jsyaml from 'js-yaml'

defineOptions({
    name: 'UploadPrompt',
})

const props = defineProps({
    modelValue: {
        type: Array,
        default: () => []
    },
    limit: {
        type: Number,
        default: 1
    },
    accept: {
        type: String,
        default: '.yaml'
    },
})

const fileList = ref([])
const emits = defineEmits(['update:modelValue', 'on-success', 'on-error'])

const onChange = (file) => {
    const reader = new FileReader()
    reader.onload = (event) => {
        try {
            const yamlContent = event.target.result
            const jsonObject = jsyaml.load(yamlContent)
            const jsonString = JSON.stringify(jsonObject, null, 2)
            console.log(jsonString)
            console.log(JSON.parse(jsonString))
            emits('update:modelValue', [jsonString])  // 将 jsonString 包装成数组传递
            emits('on-success', jsonString)
        } catch (error) {
            console.error('Error in ReadYaml:', error)
            ElMessage({ type: 'error', message: error.message })
            emits('on-error', error)
        }
    }
    reader.readAsText(file.raw)
}

watch(
    () => props.modelValue,
    value => {
        fileList.value = value.length > 0 ? [{ name: 'uploaded.yaml', url: '' }] : []
    },
    { immediate: true }
)
</script>
