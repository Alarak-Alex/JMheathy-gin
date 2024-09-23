<template>
    <div class="array-ctrl-container">
        <draggable v-model="modelValue" :disabled="!editable" handle=".drag-handle" class="w-full">
            <template v-slot:item="{ element, index }">
                <div class="flex items-center gap-2 w-full">
                    <el-tag v-if="!isEditing(index) && !inDrawer" :key="element" :closable="editable"
                        :disable-transitions="false" @close="handleClose(element)" @click="startEditing(index)"
                        :class="['drag-handle', { 'expanded-tag': isEditing(index) }]">
                        {{ element }}
                    </el-tag>
                    <el-input v-else v-model="editValue" size="small" class="input-class"
                        @keyup.enter="confirmEdit(index)" @blur="confirmEdit(index)" />
                    <span v-if="editable" class="drag-icon">☰</span>
                </div>
            </template>
        </draggable>
        <template v-if="editable">
            <el-input autosize v-if="inputVisible" ref="InputRef" v-model="inputValue" class="input-class" size="small"
                @keyup.enter="handleInputConfirm" @blur="handleInputConfirm" />
            <el-button v-else class="button-new-tag" size="small" @click="showInput">
                + 新增
            </el-button>
        </template>
    </div>
</template>

<script setup>
defineOptions({
    name: 'ArrayCtrl',
})

import { nextTick, ref } from 'vue'
import draggable from 'vuedraggable'

const inputValue = ref('')
const inputVisible = ref(false)
const InputRef = ref(null)
const editValue = ref('')

const modelValue = defineModel()

const editingIndex = ref(null)

defineProps({
    editable: {
        type: Boolean,
        default: () => false
    },
    inDrawer: {
        type: Boolean,
        default: () => false
    }
})

const handleClose = (tag) => {
    modelValue.value.splice(modelValue.value.indexOf(tag), 1)
}

const showInput = () => {
    inputVisible.value = true
    nextTick(() => {
        InputRef.value?.input?.focus()
    })
}

const handleInputConfirm = () => {
    if (inputValue.value) {
        modelValue.value.push(inputValue.value)
    }
    inputVisible.value = false
    inputValue.value = ''
}

const startEditing = (index) => {
    editingIndex.value = index
    editValue.value = modelValue.value[index]
    nextTick(() => {
        InputRef.value?.input?.focus()
    })
}

const confirmEdit = (index) => {
    if (editValue.value) {
        modelValue.value[index] = editValue.value
    }
    editingIndex.value = null
}

const isEditing = (index) => {
    return editingIndex.value === index
}
</script>

<style scoped>
.array-ctrl-container {
    display: flex;
    flex-direction: column;
    width: 100%;
}

.flex {
    display: flex;
}

.flex-col {
    flex-direction: column;
}

.items-start {
    align-items: flex-start;
}

.gap-2 {
    gap: 0.5rem;
    /* Adjust the gap size as needed */
}

.w-full {
    width: 100%;
}

.input-class .el-input__inner {
    white-space: normal;
    word-break: break-word;
    max-width: 100%;
    /* Ensure the input does not exceed its container */
}

.drag-handle {
    cursor: grab;
}

.drag-icon {
    cursor: grab;
}

.expanded-tag {
    height: calc(1em * 7);
    display: flex;
    align-items: center;
}

.el-input {
    width: 100%;
    /* Ensure the input takes the full width of its container */
}

.el-input .el-input__inner {
    resize: none;
    /* Disable manual resizing */
    overflow: auto;
    /* Enable scrollbars */
}
</style>