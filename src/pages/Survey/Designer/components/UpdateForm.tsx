import {
  ProFormText,
  ProForm, ProFormSelect, ProFormRadio,
} from '@ant-design/pro-components';

import { Button, Divider, message, Modal } from 'antd';
import React, { useState } from 'react';
import { DeleteOutlined,MenuOutlined, PlusOutlined } from '@ant-design/icons';
import { createQuestion, updateQuestion, deleteQuestion } from '@/services/survey';

export type FormValueType = {
  target?: string;
  template?: string;
  type?: string;
  time?: string;
  frequency?: string;
} & Partial<API.Survey>;

export type UpdateFormProps = {
  onCancel: (flag?: boolean, formVals?: FormValueType) => void;
  onSubmit: (values: FormValueType) => Promise<void>;
  updateModalOpen: boolean;
  values: Partial<API.Survey>;
};




const UpdateForm: React.FC<UpdateFormProps> = (props) => {

  console.log(props)











  return (
    <>
    {/* 问题编辑模态框 */}
  <Modal
    title={editingQuestion ? '编辑问题' : '添加问题'}
    open={visible}
    onCancel={() => setVisible(false)}
    footer={null}
    width={800}
  >
    <ProForm
      form={form}
      layout="vertical"
      initialValues={editingQuestion || { required: true, sort_order: questions.length + 1 }}
    >
      <ProFormText
        name="content"
        label="问题内容"
        rules={[{ required: true, message: '请输入问题内容' }]}
        rows={3}
      />

      <ProFormSelect
        name="question_type"
        label="问题类型"
        onChange={(value) => setQuestionType(value as QuestionType)}
        options={[
          { label: '单选题', value: 'single_choice' },
          { label: '多选题', value: 'multiple_choice' },
          { label: '文本题', value: 'text' },
          { label: '数字题', value: 'number' },
          { label: '日期题', value: 'date' },
          { label: '矩阵单选题', value: 'matrix_single' },
        ]}
      />

      <ProFormRadio.Group
        name="required"
        label="是否必填"
        options={[
          { label: '是', value: true },
          { label: '否', value: false },
        ]}
      />

      <ProFormText
        name="sort_order"
        label="排序"
        type="number"
        placeholder="请输入排序号"
      />

      {/* 选项设置（针对单选题和多选题） */}
      {['single_choice', 'multiple_choice'].includes(questionType) && (
        <div>
          <h3 className="font-medium mb-2">选项设置</h3>
          {options.map((option, index) => (
            <div key={index} className="flex items-center mb-2">
              <ProFormText
                value={option}
                onChange={(value) => {
                  const newOptions = [...options];
                  newOptions[index] = value;
                  setOptions(newOptions);
                }}
                placeholder={`选项 ${index + 1}`}
                style={{ marginRight: 8, flex: 1 }}
              />
              <Button
                type="danger"
                icon={<DeleteOutlined />}
                onClick={() => {
                  const newOptions = [...options];
                  newOptions.splice(index, 1);
                  setOptions(newOptions);
                }}
                size="small"
              />
            </div>
          ))}
          <Button
            type="dashed"
            onClick={() => setOptions([...options, ''])}
            style={{ width: '100%' }}
          >
            <PlusOutlined /> 添加选项
          </Button>
        </div>
      )}

      {/* 矩阵设置（针对矩阵单选题） */}
      {questionType === 'matrix_single' && (
        <div>
          <h3 className="font-medium mb-3">矩阵设置</h3>

          <div className="mb-6">
            <h4 className="font-medium mb-2">行设置</h4>
            {matrixRows.map((row, index) => (
              <div key={index} className="flex items-center mb-2">
                <ProFormText
                  value={row}
                  onChange={(value) => {
                    const newRows = [...matrixRows];
                    newRows[index] = value;
                    setMatrixRows(newRows);
                  }}
                  placeholder={`行 ${index + 1}`}
                  style={{ marginRight: 8, flex: 1 }}
                />
                <Button
                  type="danger"
                  icon={<DeleteOutlined />}
                  onClick={() => {
                    const newRows = [...matrixRows];
                    newRows.splice(index, 1);
                    setMatrixRows(newRows);
                  }}
                  size="small"
                />
              </div>
            ))}
            <Button
              type="dashed"
              onClick={() => setMatrixRows([...matrixRows, ''])}
              style={{ width: '100%' }}
            >
              <PlusOutlined /> 添加行
            </Button>
          </div>

          <div>
            <h4 className="font-medium mb-2">列设置</h4>
            {matrixColumns.map((column, index) => (
              <div key={index} className="flex items-center mb-2">
                <ProFormText
                  value={column}
                  onChange={(value) => {
                    const newColumns = [...matrixColumns];
                    newColumns[index] = value;
                    setMatrixColumns(newColumns);
                  }}
                  placeholder={`列 ${index + 1}`}
                  style={{ marginRight: 8, flex: 1 }}
                />
                <Button
                  type="danger"
                  icon={<DeleteOutlined />}
                  onClick={() => {
                    const newColumns = [...matrixColumns];
                    newColumns.splice(index, 1);
                    setMatrixColumns(newColumns);
                  }}
                  size="small"
                />
              </div>
            ))}
            <Button
              type="dashed"
              onClick={() => setMatrixColumns([...matrixColumns, ''])}
              style={{ width: '100%' }}
            >
              <PlusOutlined /> 添加列
            </Button>
          </div>
        </div>
      )}

      <Divider />

      <div className="flex justify-end mt-4">
        <Button onClick={() => setVisible(false)} style={{ marginRight: 8 }}>
          取消
        </Button>
        <Button type="primary" onClick={handleSaveQuestion}>
          保存
        </Button>
      </div>
    </ProForm>
  </Modal>
    </>
  );
};

export default UpdateForm ;
