<template>
  <div class="mb-1">
    <label :for="name" class="form-label">{{label}}</label>
    <textarea
        :name="name"
        :placeholder="placeholder"
        :readonly="readonly"
        :required="required"
        :min="min"
        :max="max"
        :value="modelValue"
        :autocomplete="name + '-new'"
        @input="$emit('update:modelValue', $event.target.value)"
        class="form-control">
    </textarea>
  </div>
</template>

<script>
export default {
  name: 'TextArea',
  props: {
    name: String,
    label: String,
    placeholder: String,
    required: String,
    min: String,
    max: String,
    modelValue: String,
    readonly: Boolean,
  },
}
</script>

<style scoped>


</style>