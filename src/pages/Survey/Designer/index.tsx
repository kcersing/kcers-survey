import React, { useState, useRef,useEffect } from 'react';
import { useParams } from "react-router"
import { FormattedMessage } from '@umijs/max';
import {
  type ActionType,
  ProColumns,
  ProForm,
  ProFormRadio,
  ProFormSelect,
  ProFormText,
  ProFormDateTimePicker,
  ProFormTextArea,
  ProCard,
  ProFormDigit,
  DragSortTable,

  ProFormGroup,
  ProFormList,

} from '@ant-design/pro-components';
import { Button, message, Modal, Divider ,Segmented } from 'antd';
import { listQuestion, getSurvey,createQuestion, updateQuestion, deleteQuestion } from '@/services/ant-design-pro/survey';

import { DeleteOutlined, MenuOutlined, PlusOutlined ,CloseCircleOutlined, SmileOutlined, SnippetsOutlined,CloseOutlined} from '@ant-design/icons';

type QuestionType = 'single_choice' | 'multiple_choice' | 'text' | 'number' | 'date' | 'matrix_single';


const Designer =  () => {

  const actionRef = useRef<ActionType>();


  const [questions, setQuestions] = useState<API.Questions[]>([]);
  const [survey, setSurvey] = useState<any>({});
  const [loading, setLoading] = useState(false);
  const [visible, setVisible] = useState(false);
  const [editingQuestion, setEditingQuestion] = useState<API.Questions | null>(null);
  const [questionType, setQuestionType] = useState<QuestionType>('single_choice');
  const [options, setOptions] = useState<API.Options[]>([]);
  const [matrixRows, setMatrixRows] = useState<string[]>([]);
  const [matrixColumns, setMatrixColumns] = useState<string[]>([]);
  const [form] = ProForm.useForm();


  let params = useParams();
  const surveyId:number =parseInt( params.id)
  useEffect(() => {
    loadSurveyAndQuestions();
  }, []);

  const loadSurveyAndQuestions = async () => {
    try {
      setLoading(true);

      // 并行加载问卷和问题
      const [surveyData, questionsData] = await Promise.all([
        getSurvey({'id':surveyId}),
        listQuestion({'surveyId':surveyId}),
      ]);

      setSurvey(surveyData.data);
      setQuestions(questionsData.data);

    } catch (error) {
      message.error('加载问卷数据失败');
    } finally {
      setLoading(false);
    }
  };



  const handleDragSortEnd = (
    beforeIndex: number,
    afterIndex: number,
    newDataSource: any,
  ) => {
    console.log('beforeIndex', beforeIndex);
    console.log('afterIndex', afterIndex);
    console.log('排序后的数据', newDataSource);
    setQuestions(newDataSource)
    // 请求成功之后刷新列表
    actionRef.current?.reload();
    message.success('修改列表排序成功');
  };

  const dragHandleRender = (rowData: any, idx: any) => (
    <>
      <MenuOutlined style={{ cursor: 'grab', color: 'gold' }} />
      &nbsp;{idx + 1} - {rowData.name}
    </>
  );

  const columns: ProColumns[] = [
    {
      title: '排序',
      dataIndex: 'sort',
      width: 60,
      className: 'drag-visible',
    },
    {
      title: '问题内容',
      dataIndex: 'content',
      className: 'drag-visible',
    },
    {
      title: '问题类型',
      dataIndex: 'type',
      valueEnum: {
        single_choice: '单选题',
        multiple_choice: '多选题',
        text: '文本题',
        number: '数字题',
        date: '日期题',
        matrix_single: '矩阵单选题',
      },
    },
    {
      title: '必填',
      dataIndex: 'required',
      valueType: 'select',
      valueEnum: {
        true: '是',
        false: '否',
      },
    },
    {
      title: '操作',
      key: 'option',
      valueType: 'option',
      render: (_, record) => [
        <a key="ex" onClick={() => handleEditQuestion(record)}>编辑</a>,
        <a key="de" onClick={() => handleDeleteQuestion(record.id)} >删除</a>,
      ],
    },
  ];



  const handleAddQuestion = () => {
    // 重置表单和状态
    form.resetFields();
    setEditingQuestion(null);
    setQuestionType('single_choice');
    setOptions([]);
    setMatrixRows([]);
    setMatrixColumns([]);
    setVisible(true);
  };

  const handleSaveQuestion = async () => {
    try {
      const values = await form.validateFields();
console.log(values);
      const questionData = {
        ...values,
        survey_id: surveyId,
        question_type: questionType,
      };

      if (questionType === 'single_choice' || questionType === 'multiple_choice') {
        questionData.options = options;
      } else if (questionType === 'matrix_single') {
        // questionData.matrix_rows = matrixRows.filter(row => row.trim()).join(',');
        // questionData.matrix_columns = matrixColumns.filter(column => column.trim()).join(',');
      }
      console.log(questionData)
      if (editingQuestion) {
        // 更新问题
        await updateQuestion(editingQuestion.id, questionData);
        message.success('问题更新成功');
      } else {
        // 创建新问题
        questionData.sort = questions.length + 1;
        await createQuestion(questionData);
        message.success('问题创建成功');
      }

      loadSurveyAndQuestions();
      setVisible(false);
    } catch (errorInfo) {
      console.log('表单验证失败:', errorInfo);
      message.error('表单验证失败，请检查输入');
    }
  };

  const handleEditQuestion = (question: API.Questions) => {
    // 填充表单和状态

    form.setFieldsValue({
      content: question.content,
      question_type: question.type,
      required: question.required,
      sort: question.sort,
    });

    setEditingQuestion(question);
    setQuestionType(question.type);

    if (question.type === 'single_choice' || question.type === 'multiple_choice') {
       setOptions(question.options);
    } else if (question.type === 'matrix_single') {
      // setMatrixRows(question.matrixRows?.split(',') || []);
      // setMatrixColumns(question.matrixColumns?.split(',') || []);
    }

    setVisible(true);
  };

  const handleDeleteQuestion = async (id: number) => {
    try {
      await deleteQuestion({'id':id});
      message.success('问题删除成功');
    } catch (error) {
      message.error('问题删除失败');
    }
  };



  return (
<>
    <ProCard
      title={`问卷设计: ${survey.title || '加载中...'}`}
      extra={[
        <Button key="preview" href={`/survey/${surveyId}/respond`}>
          预览问卷
        </Button>,
        <Button key="add" type="primary" onClick={handleAddQuestion}>
          <PlusOutlined /> 添加问题
        </Button>,
      ]}
    >

      <DragSortTable
        actionRef={actionRef}
        headerTitle="拖拽排序(默认把手)"
        columns={columns}
        rowKey="id"
        search={false}
        pagination={false}
        loading={loading}
        dataSource={questions}
        dragSortKey="sort"

        dragSortHandlerRender={dragHandleRender}
        onDragSortEnd={handleDragSortEnd}
      />
    </ProCard>


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
      initialValues={editingQuestion || { required: true, sort: questions.length + 1 }}
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

      <ProFormDigit
        name="sort"
        label="排序"
        min={1}
        max={10}
        fieldProps={{ precision: 0 }}
        placeholder="请输入排序号"
      />

      {/* 选项设置（针对单选题和多选题） */}
      {['single_choice', 'multiple_choice'].includes(questionType) && (
        <div>
          <h3 className="font-medium mb-2">选项设置</h3>
          <ProFormList
            copyIconProps={{ Icon: SnippetsOutlined, }}
            initialValue={options.option}
            deleteIconProps={{ Icon: CloseOutlined, }}
            ksy="options"
            name="options"
          >
            <ProFormText hidden={true}   name="serial" label="序号" />
            <ProFormText  name="content" label="选项" />
          </ProFormList>

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
export default Designer;
