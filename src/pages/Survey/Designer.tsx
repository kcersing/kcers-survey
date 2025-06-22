import React, { useState, useRef,useEffect } from 'react';

import { ProForm, ProFormText, ProFormSelect, ProFormRadio, ProFormGroup,DragSortTable, } from '@ant-design/pro-components';
import { PlusOutlined, DeleteOutlined, EditOutlined, DragOutlined } from '@ant-design/icons';
import { ProTable } from '@ant-design/pro-components';
import {
    listQuestion,
    createQuestion,
    updateQuestion,
    deleteQuestion,
    getSurvey,
    listSurvey
} from '@/services/ant-design-pro/survey';
import { MenuOutlined } from '@ant-design/icons';
import type { ActionType, ProColumns } from '@ant-design/pro-components';

import { SortableContainer, SortableElement, arrayMove } from 'react-sortable-hoc';



import {Button, Modal, message, Card, Space, Divider  } from 'antd';



type QuestionType = 'single_choice' | 'multiple_choice' | 'text' | 'number' | 'date' | 'matrix_single';


// 拖拽排序项
const SortableItem = SortableElement(({ question, onEdit, onDelete }) => (
    <tr>
      <td className="cursor-move"><DragOutlined /></td>
      <td>{question.sort}</td>
      <td>{question.content}</td>
      <td>
        {(() => {
          const typeMap = {
            single_choice: '单选题',
            multiple_choice: '多选题',
            text: '文本题',
            number: '数字题',
            date: '日期题',
            matrix_single: '矩阵单选题',
          };
          return typeMap[question.type] || question.type;
        })()}
      </td>
      <td>{question.required ? '是' : '否'}</td>
      <td>
        <Space>
          <Button icon={<EditOutlined />} onClick={() => onEdit(question)} size="small">
            编辑
          </Button>
          <Button icon={<DeleteOutlined />} danger onClick={() => onDelete(question.id)} size="small">
            删除
          </Button>
        </Space>
      </td>
    </tr>
));



const SurveyDesigner: React.FC<{ match }> = ({ match }) => {

  const surveyId = 1;
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
      setMatrixRows(question.matrixRows?.split(',') || []);
      setMatrixColumns(question.matrixColumns?.split(',') || []);
    }

    setVisible(true);
  };

  const handleDeleteQuestion = async (id: number) => {
    try {
      await deleteQuestion({'id':id});
      message.success('问题删除成功');
      loadSurveyAndQuestions();
    } catch (error) {
      message.error('问题删除失败');
    }
  };

  const handleSaveQuestion = async () => {
    try {
      const values = await form.validateFields();

      const questionData = {
        ...values,
        survey_id: surveyId,
        question_type: questionType,
      };

      if (questionType === 'single_choice' || questionType === 'multiple_choice') {
        questionData.options = options.filter(option => option.trim()).join(',');
      } else if (questionType === 'matrix_single') {
        questionData.matrix_rows = matrixRows.filter(row => row.trim()).join(',');
        questionData.matrix_columns = matrixColumns.filter(column => column.trim()).join(',');
      }

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

  const onSortEnd = ({ oldIndex, newIndex }: { oldIndex: number; newIndex: number }) => {
    if (oldIndex !== newIndex) {
      const newQuestions = arrayMove(questions, oldIndex, newIndex);

      // 更新排序
      const updatePromises = newQuestions.map((question, index) => {
        if (question.sort !== index + 1) {
          return updateQuestion(question.id, { sort_order: index + 1 });
        }
        return Promise.resolve();
      });

      Promise.all(updatePromises)
          .then(() => {
            message.success('排序更新成功');
            loadSurveyAndQuestions();
          })
          .catch(() => {
            message.error('排序更新失败');
          });
    }
  };
    const actionRef = useRef<ActionType>();

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
                <a key="de" onClick={() => handleDeleteQuestion(record.id)} danger>删除</a>,
            ],
        },
    ];
    const handleDragSortEnd1 = (
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

  return (
      <div>
        <Card
            title={`问卷设计: ${survey.title || '加载中...'}`}
            bordered={false}
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
                onDragSortEnd={handleDragSortEnd1}
            />



        </Card>

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
      </div>
  );
};

export default SurveyDesigner;
