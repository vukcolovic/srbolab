<template>
    <div class="mb-1">
        <label :for="name" class="form-label">{{label}}</label>
        <input
            :type="type"
            :name="name"
            :placeholder="placeholder"
            :readonly="readonly"
            :required="required"
            :step= "step"
            :min="min"
            :max="max"
            :value="modelValue"
            :autocomplete="name + '-new'"
            @input="$emit('update:modelValue', $event.target.value)"
            :ref="refVal"
            class="form-control">
    </div>
</template>

<script>
export default {
    name: 'TextInput',
    props: {
        name: String,
        type: String,
        label: String,
        placeholder: String,
        required: Boolean,
        min: String,
        max: String,
        step: String,
        modelValue: String,
        readonly: Boolean,
        refVal: String,
    },
}
</script>

<style scoped>


</style>