import React, { useState, useRef,useEffect } from 'react';
import { useParams } from "react-router"
import {
  type ActionType,
  ProColumns,
  ProForm,
  ProFormRadio,
  ProFormSelect,
  ProFormText,
  ProCard,
  ProFormDigit,
  ProTable,
  ProFormGroup,
  ProFormList,
  ProFormTreeSelect,

} from '@ant-design/pro-components';
import { Button, message, Modal, Divider  } from 'antd';
import {
  listQuestion,
  getSurvey,
  createQuestion,
  updateQuestion,
  deleteQuestion,
  treeQuestion,
} from '@/services/ant-design-pro/survey';

import { PlusOutlined , SnippetsOutlined,CloseOutlined} from '@ant-design/icons';

type QuestionType = 'h2' | 'h3' | 'page' | 'single_choice' | 'multiple_choice' | 'dropdown' | 'text' | 'number' | 'slider' | 'date' | 'datetime' | 'rate' | 'nps' | 'ranking' | 'matrix' | 'image' | 'file' | 'video' | 'audio' | 'signature';


const Designer =  () => {

  const actionRef = useRef<ActionType>();
  const [questions, setQuestions] = useState<API.Questions[]>([]);
  const [survey, setSurvey] = useState<API.Survey>({});
  const [loading, setLoading] = useState(false);
  const [visible, setVisible] = useState(false);
  const [editingQuestion, setEditingQuestion] = useState<API.Questions | null>(null);
  const [questionType, setQuestionType] = useState<QuestionType>('single_choice');
  const [options, setOptions] = useState<API.Options[]>([]);
  const [form] = ProForm.useForm();


  let params = useParams();
  const surveyId =parseInt(params.id)
  useEffect(() => {
    loadSurveyAndQuestions();
  }, []);

  const loadSurveyAndQuestions = async () => {
    try {
      setLoading(true);

      // 并行加载问卷和问题
      const [surveyData, questionsData] = await Promise.all([
        getSurvey({id:surveyId}),
        listQuestion({surveyId:surveyId}),
      ]);

      setSurvey(surveyData.data);
      setQuestions(questionsData.data);

    } catch (error) {
      message.error('加载问卷数据失败');
    } finally {
      setLoading(false);
    }
  };


  const columns: ProColumns<API.Questions>[] = [
    {
      title: '排序',
      dataIndex: 'sort',
      width: 120,
      className: 'drag-visible',
    },

    {
      title: '编号',
      dataIndex: 'serial',
      width: 120,
      className: 'drag-visible',
    },

    {
      title: '问题内容',
      dataIndex: 'content',
      className: 'drag-visible',
      ellipsis: true,
    },
    {
      title: '问题类型',
      dataIndex: 'type',
      valueEnum: {
        h2:'标题1',
        h3:'标题2',
        page:'单页',
        single_choice: '单选题',
        multiple_choice: '多选题',
        dropdown: '下拉题',
        text: '文本题',
        number: '数字题',
        slider: '滑动条',
        date: '日期题',
        datetime: '日期时间',
        rate:'评分',
        nps:'NPS评分',
        ranking: '排序题',
        matrix: '矩阵题',
        image: '图片上传',
        file: '文件上传',
        video: '视频上传',
        audio: '音频上传',
        signature: '签名',
      },
    },
    {
      title: '必填',
      dataIndex: 'required',
      valueType: 'select',
      valueEnum: {
        1: '是',
        2: '否',
      },
    },
    {
      title: '操作',
      key: 'option',
      valueType: 'option',
      render: (_, record) => [
        <a key="ex" onClick={() =>  handleEditQuestion(record)}>编辑</a>,
        <a key="de" onClick={() => handleDeleteQuestion(record.id)} >删除</a>,
      ],
    },
  ];


  const handleChange = (value: string[]) => {
  };

  const handleAddQuestion = () => {
    // 重置表单和状态
    form.resetFields();

    setEditingQuestion(null);
    setQuestionType('single_choice');
    setOptions([]);
    setVisible(true);
  };

  const handleSaveQuestion = async () => {
    try {
      const values = await form.validateFields();

      let questionData = {
        ...values,
        surveyId: surveyId,
        questionType: questionType,
        type: questionType,
        parentId:parseInt(values.parentId),
        required:parseInt(values.required),
      };

      if (questionType === 'single_choice' || questionType === 'multiple_choice') {
        questionData.options = values.options;
      }

      if (editingQuestion) {
        // 更新问题
        await updateQuestion({'id':editingQuestion.id, ...questionData});
        message.success('问题更新成功');
      } else {
        // 创建新问题
        questionData.sort =  questions ? questions.length + 1 : 0;
        await createQuestion(questionData);
        message.success('问题创建成功');
      }

      loadSurveyAndQuestions();
      setVisible(false);
    } catch (errorInfo) {
      message.error('表单验证失败，请检查输入');
    }
  };

  const handleEditQuestion = (question: API.Questions) => {
    // 填充表单和状态
    handleAddQuestion()
    form.setFieldsValue({
      content: question.content,
      questionType: question.type,
      required: question.required,
      sort: question.sort,
      jumpRules: question.jumpRules,
      parentId: question.parentId,
    });

    setEditingQuestion(question);
    setQuestionType(question.type);

    if (question.type === 'single_choice' || question.type === 'multiple_choice') {
       setOptions(question.options);
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
        <ProTable<API.Questions>
          actionRef={actionRef}
          columns={columns}
          rowKey="id"
          search={false}
          pagination={false}
          loading={loading}
          dataSource={questions}

          columnWidth={12}
          defaultExpandAllRows={true}
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
          initialValues={editingQuestion || { required: 1,  questionType:'single_choice', sort: questions ? questions.length + 1 : 0 }}
          submitter={{
            // 配置按钮文本
            searchConfig: {
              resetText: '重置',
              submitText: '提交',
            },
            // 配置按钮的属性
            resetButtonProps: {
              style: {
                // 隐藏重置按钮
                display: 'none',
              },
            },
            submitButtonProps: {},

            // 完全自定义整个区域
            render: (props, doms) => {
              return (
                <div className="flex justify-end mt-4">
                  <Button onClick={() => setVisible(false)} style={{ marginRight: 8 }}>
                    取消
                  </Button>
                  <Button type="primary" onClick={handleSaveQuestion}>
                    保存
                  </Button>
                </div>
              );
            },
          }}
        >
          <ProFormTreeSelect
            name="parentId"
            label="上级问题"
            params={{}}

            request={async ({ keyWords }) => {
              const questionAll = await treeQuestion({ surveyId: surveyId, keywords: keyWords });

              return questionAll.data;

            }}
            fieldNames={{ label: 'title', value: 'value', children: 'children' }}

            style={{ width: '100%' }}
            placeholder="Please select"
            onChange={handleChange}
            showSearch={false}

          />
          <ProFormText
              name="serial"
              label="编号"
          />
          <ProFormText
            name="content"
            label="问题内容"
            rules={[{ required: true, message: '请输入问题内容' }]}
            rows={3}
          />

          <ProFormSelect
            name="questionType"
            label="问题类型"
            onChange={(value) => setQuestionType(value as QuestionType)}
            options={[
              { label: '标题1', value: 'h2' },
              { label: '标题2', value: 'h3' },
              { label: '单页', value: 'page' },
              { label: '单选题', value: 'single_choice' },
              { label: '多选题', value: 'multiple_choice' },
              { label: '下拉题', value: 'dropdown' },
              { label: '文本题', value: 'text' },
              { label: '数字题', value: 'number' },
              { label: '滑动条', value: 'slider' },
              { label: '日期题', value: 'date' },
              { label: '日期时间', value: 'datetime' },
              { label: '评分', value: 'rate' },
              { label: 'NPS评分', value: 'nps' },
              { label: '排序题', value: 'ranking' },
              { label: '矩阵题', value: 'matrix' },
              { label: '图片上传', value: 'image' },
              { label: '文件上传', value: 'file' },
              { label: '视频上传', value: 'video' },
              { label: '音频上传', value: 'audio' },
              { label: '签名', value: 'signature' },
            ]}
          />

          <ProFormRadio.Group
            name="required"
            label="是否必填"
            options={[
              { label: '是', value: 1 },
              { label: '否', value: 2 },
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


          {questionType === 'single_choice' && (
            <>
              <ProFormList
                label="选项设置"
                copyIconProps={{ Icon: SnippetsOutlined }}
                initialValue={options}
                deleteIconProps={{ Icon: CloseOutlined }}
                name="options"
              >
                  <ProFormGroup key="group">
                  <ProFormDigit name="serial" label="序号"/>
                  <ProFormText  name="content" label="选项" />

                  <ProFormRadio.Group
                      name="inputs"
                      label="是否可填写"
                      value={1}
                      options={[
                        { label: '否', value: 1 },
                        { label: '是', value: 2 },
                      ]}
                  />
                </ProFormGroup>
              </ProFormList>
            </>
          )}
          {questionType === 'multiple_choice' && (
            <>
              <ProFormList
                label="选项设置"
                copyIconProps={{ Icon: SnippetsOutlined }}
                initialValue={options}
                deleteIconProps={{ Icon: CloseOutlined }}
                name="options"
              >
                <ProFormGroup key="group">
                  <ProFormDigit name="serial" label="序号"/>
                  <ProFormText  name="content" label="选项" />

                  <ProFormRadio.Group
                      name="inputs"
                      label="是否可填写"
                      value={1}
                      options={[
                        { label: '否', value: 1 },
                        { label: '是', value: 2 },
                      ]}
                  />
                </ProFormGroup>
              </ProFormList>
            <ProFormDigit
            name="valueNumber"
            label="最多选中"
            fieldProps={{ precision: 0 }}
            placeholder="3"
            />
            </>
          )}




          <Divider />
        </ProForm>
      </Modal>
    </>
  );
};
export default Designer;
