import { ProFormDateTimePicker, ProFormText, ProFormTextArea, ProForm, } from '@ant-design/pro-components';
import { Modal } from 'antd';
import React from 'react';
const UpdateForm = (props) => {
    return (<ProForm contentRender={(items, submitter, form) => {
            return (<Modal width={640} destroyOnClose title='编辑问卷' open={props.updateModalOpen} footer={submitter} onCancel={() => {
                    props.onCancel();
                }}>
            {items}
          </Modal>);
        }} onFinish={props.onSubmit}>

        <ProFormText name="title" key="title" label='标题' initialValue={props.values.title} rules={[
            {
                required: true,
                message: '请输入问卷标题',
            },
        ]}/>


        <ProFormTextArea name="desc" key="desc" label='描述' initialValue={props.values.desc} placeholder='请输入至少五个字符' rules={[
            {
                required: true,
                message: '请输入至少五个字符的描述',
                min: 5,
            },
        ]}/>

        <ProFormDateTimePicker name="startAt" key="startAt" label='开始时间' initialValue={props.values.startAt} rules={[
            {
                required: true,
                message: '请选择开始时间',
            },
        ]}/>
        <ProFormDateTimePicker name="endAt" key="endAt" label='结束时间' initialValue={props.values.endAt} rules={[
            {
                required: true,
                message: '请选择结束时间',
            },
        ]}/>

    </ProForm>);
};
export default UpdateForm;
//# sourceMappingURL=UpdateForm.jsx.map