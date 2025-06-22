import React, { useState, useEffect } from 'react';
import { ProForm, ProFormText, ProFormSelect, ProFormDatePicker, ProFormRadio, ProFormTextArea,  } from '@ant-design/pro-components';
import { LoadingOutlined } from '@ant-design/icons';
import { fetchQuestions, submitResponse,  getSurvey } from '@/services/ant-design-pro/survey';
import {Button, Card, message, Steps, Space  } from 'antd';
const { Step } = Steps;

type QuestionType = 'single_choice' | 'multiple_choice' | 'text' | 'number' | 'date' | 'matrix_single';

interface Question {
  id: number;
  content: string;
  question_type: QuestionType;
  options?: string;
  matrix_rows?: string;
  matrix_columns?: string;
  required: boolean;
}

const SurveyRespondent: React.FC<{ match: { params: { id: string } } }> = ({ match }) => {
  const surveyId = parseInt(match.params.id, 10);
  const [questions, setQuestions] = useState<Question[]>([]);
  const [survey, setSurvey] = useState<any>({});
  const [currentStep, setCurrentStep] = useState(0);
  const [loading, setLoading] = useState(true);
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
        fetchQuestions(surveyId),
      ]);

      setSurvey(surveyData);
      setQuestions(questionsData.sort((a, b) => a.sort_order - b.sort_order));

      // 初始化表单值
      const initialValues: any = {};
      questionsData.forEach(question => {
        initialValues[`question_${question.id}`] = '';
      });
      form.setFieldsValue(initialValues);
    } catch (error) {
      message.error('加载问卷失败');
    } finally {
      setLoading(false);
    }
  };

  const handleNext = async () => {
    try {
      await form.validateFields([`question_${questions[currentStep].id}`]);
      if (currentStep < questions.length - 1) {
        setCurrentStep(currentStep + 1);
      } else {
        // 提交问卷
        handleSubmit();
      }
    } catch (errorInfo) {
      console.log('Validation Failed:', errorInfo);
    }
  };

  const handlePrev = () => {
    if (currentStep > 0) {
      setCurrentStep(currentStep - 1);
    }
  };

  const handleSubmit = async () => {
    try {
      setLoading(true);
      const values = await form.validateFields();

      // 构建提交数据
      const answers = questions.map(question => ({
        question_id: question.id,
        answer: values[`question_${question.id}`] || '',
      }));

      // 提交回答
      await submitResponse({
        survey_id: surveyId,
        respondent: localStorage.getItem('userId') || 'anonymous',
        answers,
      });

      message.success('问卷提交成功');
      setTimeout(() => {
        window.location.href = '/survey/list';
      }, 2000);
    } catch (error) {
      message.error('问卷提交失败');
    } finally {
      setLoading(false);
    }
  };

  if (loading && !questions.length) {
    return (
        <div className="flex justify-center items-center h-screen">
          <LoadingOutlined style={{ fontSize: 40 }} spin />
        </div>
    );
  }

  const currentQuestion = questions[currentStep];
  if (!currentQuestion) {
    return <div>问卷不存在或已删除</div>;
  }

  return (
      <div>
        <Card
            title={survey.title || '问卷'}
            subTitle={survey.description || ''}
            bordered={false}
        />

        <Steps current={currentStep} className="mb-8">
          {questions.map((question, index) => (
              <Step key={question.id} title={`问题 ${index + 1}`} />
          ))}
        </Steps>

        <Card title={`问题 ${currentStep + 1}/${questions.length}`} className="mb-6" bordered={false}>
          <h2 className="text-xl font-bold mb-4">{currentQuestion.content}</h2>

          <ProForm form={form} layout="vertical">
            {/* 根据问题类型渲染不同的表单组件 */}
            {(() => {
              switch (currentQuestion.question_type) {
                case 'single_choice':
                  return (
                      <ProFormSelect
                          name={`question_${currentQuestion.id}`}
                          label="选项"
                          rules={currentQuestion.required ? [{ required: true, message: '请选择一个选项' }] : []}
                          options={currentQuestion.options?.split(',').map(option => ({
                            label: option,
                            value: option,
                          })) || []}
                      />
                  );

                case 'multiple_choice':
                  return (
                      <ProFormSelect
                          name={`question_${currentQuestion.id}`}
                          label="选项"
                          mode="multiple"
                          rules={currentQuestion.required ? [{ required: true, message: '请至少选择一个选项' }] : []}
                          options={currentQuestion.options?.split(',').map(option => ({
                            label: option,
                            value: option,
                          })) || []}
                      />
                  );

                case 'text':
                  return (
                      <ProFormTextArea
                          name={`question_${currentQuestion.id}`}
                          label="回答"
                          rows={4}
                          rules={currentQuestion.required ? [{ required: true, message: '请输入内容' }] : []}
                      />
                  );

                case 'number':
                  return (
                      <ProFormText
                          name={`question_${currentQuestion.id}`}
                          label="回答"
                          type="number"
                          rules={[
                            ...(currentQuestion.required ? [{ required: true, message: '请输入数字' }] : []),
                            { type: 'number', message: '请输入有效数字' },
                          ]}
                      />
                  );

                case 'date':
                  return (
                      <ProFormDatePicker
                          name={`question_${currentQuestion.id}`}
                          label="回答"
                          rules={currentQuestion.required ? [{ required: true, message: '请选择日期' }] : []}
                      />
                  );

                case 'matrix_single':
                  const rows = currentQuestion.matrix_rows?.split(',') || [];
                  const columns = currentQuestion.matrix_columns?.split(',') || [];

                  return (
                      <ProForm.Group name={`question_${currentQuestion.id}`}>
                        <div className="overflow-x-auto">
                          <table className="min-w-full">
                            <thead>
                            <tr>
                              <th className="py-2 px-4 border">选项</th>
                              {columns.map(column => (
                                  <th key={column} className="py-2 px-4 border">{column}</th>
                              ))}
                            </tr>
                            </thead>
                            <tbody>
                            {rows.map(row => (
                                <tr key={row}>
                                  <td className="py-2 px-4 border">{row}</td>
                                  {columns.map(column => (
                                      <td key={column} className="py-2 px-4 border text-center">
                                        <ProFormRadio
                                            value={`${row}|${column}`}
                                            onChange={(e) => {
                                              form.setFieldsValue({
                                                [`question_${currentQuestion.id}`]: e.target.value,
                                              });
                                            }}
                                        />
                                      </td>
                                  ))}
                                </tr>
                            ))}
                            </tbody>
                          </table>
                        </div>
                      </ProForm.Group>
                  );

                default:
                  return <div>不支持的问题类型</div>;
              }
            })()}
          </ProForm>
        </Card>

        <div className="flex justify-between">
          {currentStep > 0 && (
              <Button onClick={handlePrev} size="large">
                上一题
              </Button>
          )}

          <Button
              type="primary"
              onClick={handleNext}
              size="large"
              loading={loading}
          >
            {currentStep < questions.length - 1 ? '下一题' : '提交问卷'}
          </Button>
        </div>
      </div>
  );
};

export default SurveyRespondent;
