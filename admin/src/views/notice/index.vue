<template>
  <div style="margin-top: 10px">
    <div
      style="
        width: 97%;
        border-top: 1px gray dashed;
        border-bottom: 1px gray dashed;
        padding: 20px;
        margin: 0 auto;
      "
      v-for="(item, index) in list"
      :key="index"
    >
      <span style="color: blue">{{ item.title }}</span>
      <br />
      <span style="font-size: 8px; color: green">{{ item.time }}</span>
      <br />
      <span style="font-size: 8px">{{ item.desc }}</span>
      <br />
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      list: [
        {
          title: "从8.3版开始优惠赠送时长改为赠送等额金币",
          time: "2019-11-13",
          desc: "金币永不过期，长期有效。",
        },
        {
          title: "从8.3版开始优惠赠送时长改为赠送等额金币",
          time: "2019-11-13",
          desc: "金币永不过期，长期有效。",
        },
        {
          title: "从8.3版开始优惠赠送时长改为赠送等额金币",
          time: "2019-11-13",
          desc: "金币永不过期，长期有效。",
        },
        {
          title: "从8.3版开始优惠赠送时长改为赠送等额金币",
          time: "2019-11-13",
          desc: "金币永不过期，长期有效。",
        },
        {
          title: "从8.3版开始优惠赠送时长改为赠送等额金币",
          time: "2019-11-13",
          desc: "金币永不过期，长期有效。",
        },
        {
          title: "从8.3版开始优惠赠送时长改为赠送等额金币",
          time: "2019-11-13",
          desc: "金币永不过期，长期有效。",
        },
      ],
    };
  },
  methods: {},
};
</script>

<style></style>
