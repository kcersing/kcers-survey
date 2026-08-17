import { Button, Card, Checkbox, DatePicker, Descriptions, Divider, Image, Input, InputNumber, Modal, Radio, Select, Steps, message } from 'antd';
import { EditOutlined } from '@ant-design/icons';
import { PageContainer } from '@ant-design/pro-components';
import { FC, useEffect, useState } from 'react';
import React from 'react';

import { getResponse, getResponseAnswers, listQuestion, interviewerResponseUpdate, createRespondent } from '@/services/ant-design-pro/survey';
import { queryProvince, queryCity } from '@/services/ant-design-pro/api';
import useStyles from './style.style';
import { useParams, useSearchParams } from '@@/exports';
import dayjs from 'dayjs';

const { TextArea } = Input;

interface Opt {
  content?: string;
  serial?: string;
  inputs?: number; // 1=不可填, 2=可填写(其他)
}

const qTypeLabel = (t?: string) => {
  const map: Record<string, string> = {
    single_choice: '单选题', multiple_choice: '多选题', dropdown: '下拉题',
    text: '文本题', input: '输入题', number: '数字题', slider: '滑动条',
    rate: '评分题', nps: 'NPS题', date: '日期题', datetime: '日期时间题',
    ranking: '排序题', matrix: '矩阵题', image: '图片题', file: '文件题',
    video: '视频题', audio: '音频题', signature: '签名题', address: '地址题',
    page: '分页', h2: '标题', h3: '副标题',
  };
  return map[t || ''] || t || '未知';
};

interface QuestionInfo {
  id?: number;
  type?: string;
  options?: Opt[];
  valueNumber?: number; // 最多选几项
}

const Basic: FC = () => {
  const { styles } = useStyles();
  const { id, sn } = useParams();
  const [searchParams] = useSearchParams();

  const [response, setResponse] = useState<API.Response>({});
  const [responseAnswers, setResponseAnswers] = useState<API.ResponseAnswers>([]);
  const [questionMap, setQuestionMap] = useState<Record<number, QuestionInfo>>({});

  const [editModalVisible, setEditModalVisible] = useState(false);
  const [editingAnswer, setEditingAnswer] = useState<any>(null);
  const [editSelected, setEditSelected] = useState<string>('');
  const [editMultiArr, setEditMultiArr] = useState<string[]>([]);
  const [editOtherText, setEditOtherText] = useState('');
  const [editNumValue, setEditNumValue] = useState<number | null>(null);

  const [addrModalVisible, setAddrModalVisible] = useState(false);
  const [addrProvince, setAddrProvince] = useState<{ label: string; value: string } | null>(null);
  const [addrCity, setAddrCity] = useState<{ label: string; value: string } | null>(null);
  const [addrDistrict, setAddrDistrict] = useState<{ label: string; value: string } | null>(null);
  const [addrTown, setAddrTown] = useState<{ label: string; value: string } | null>(null);
  const [addrDetail, setAddrDetail] = useState('');
  const [addrSaving, setAddrSaving] = useState(false);
  const [provinceOpts, setProvinceOpts] = useState<{ label: string; value: string }[]>([]);
  const [cityOpts, setCityOpts] = useState<{ label: string; value: string }[]>([]);
  const [districtOpts, setDistrictOpts] = useState<{ label: string; value: string }[]>([]);
  const [townOpts, setTownOpts] = useState<{ label: string; value: string }[]>([]);

  const surveyId = id ? parseInt(id) : 0;
  const responseSn = sn || '';

  const isInterviewer = searchParams.get('interviewer') === '1' && !!sessionStorage.getItem('interviewer_token');
  const token = sessionStorage.getItem('interviewer_token') || '';

  useEffect(() => {
    const load = async () => {
      const [responseRes, answersRes, questionsRes] = await Promise.all([
        getResponse({ sn: responseSn }),
        getResponseAnswers({ sn: responseSn }),
        listQuestion({ page: 1, surveyId, pageSize: 999 }),
      ]);
      setResponse(responseRes.data);
      setResponseAnswers(answersRes.data || []);

      const qMap: Record<number, QuestionInfo> = {};
      const flatten = (list: any[]) => {
        for (const q of list) {
          if (q.id) qMap[q.id] = q;
          if (q.children?.length) flatten(q.children);
        }
      };
      flatten(questionsRes?.data || []);
      setQuestionMap(qMap);
    };
    load();
  }, [surveyId, responseSn]);

  const handleEdit = (item: any) => {
    const qInfo = questionMap[item.surveyQuestionId];
    const qType = qInfo?.type || 'text';
    const answer = item.answer || [];
    setEditingAnswer({ ...item, qType, qInfo });

    if (['single_choice'].includes(qType)) {
      setEditSelected(answer[0] || '');
    } else if (['multiple_choice'].includes(qType)) {
      setEditMultiArr(answer);
    } else if (['dropdown'].includes(qType)) {
      setEditSelected(answer[0] || '');
    } else if (['number', 'slider', 'rate', 'nps'].includes(qType)) {
      setEditNumValue(answer[0] ? Number(answer[0]) : null);
    } else {
      setEditSelected(answer[0] || '');
    }
    setEditOtherText(item.answerText || '');
    setEditModalVisible(true);
  };

  const handleSave = async () => {
    if (!editingAnswer) return;
    const qType = editingAnswer.qType || 'text';
    let answer: string[] = [];
    let answerText = editOtherText;

    if (qType === 'single_choice') {
      answer = editSelected ? [editSelected] : [];
    } else if (qType === 'multiple_choice') {
      answer = editMultiArr;
    } else if (qType === 'dropdown') {
      answer = editSelected ? [editSelected] : [];
    } else if (qType === 'number' || qType === 'slider' || qType === 'rate' || qType === 'nps') {
      answer = editNumValue != null ? [String(editNumValue)] : [];
    } else if (qType === 'text' || qType === 'input') {
      answerText = editOtherText;
    } else {
      answer = editSelected ? [editSelected] : [];
    }

    try {
      const res = await interviewerResponseUpdate(
        { surveyId, sn: responseSn, questionId: editingAnswer.surveyQuestionId, answer, answerText, type: qType },
        token,
      );
      if (res.code === 0) {
        message.success('修改成功');
        setResponseAnswers((prev: any) =>
          (prev || []).map((a: any) =>
            a.id === editingAnswer.id ? { ...a, answer, answerText } : a,
          ),
        );
        setEditModalVisible(false);
      } else {
        message.error(res.msg || '修改失败');
      }
    } catch {
      message.error('修改失败');
    }
  };

  const handleAddrSave = async () => {
    setAddrSaving(true);
    try {
      const addrQ = Object.values(questionMap).find((q) => q.type === 'address');
      const addrQId = addrQ?.id || 0;

      const fields: { type: string; value: string }[] = [];
      if (addrProvince?.label) fields.push({ type: 'area', value: addrProvince.label });
      if (addrCity?.label) fields.push({ type: 'city', value: addrCity.label });
      if (addrDistrict?.label) fields.push({ type: 'district', value: addrDistrict.label });
      if (addrTown?.label) fields.push({ type: 'village', value: addrTown.label });
      if (addrDetail) fields.push({ type: 'address', value: addrDetail });

      await Promise.all(fields.map((f) =>
        createRespondent({ surveyId, type: f.type, questionId: addrQId, value: [f.value], sn: responseSn }),
      ));

      setResponse((prev: any) => ({
        ...prev,
        area: addrProvince?.label || prev?.area,
        city: addrCity?.label || prev?.city,
        district: addrDistrict?.label || prev?.district,
        village: addrTown?.label || prev?.village,
        address: addrDetail || prev?.address,
      }));
      message.success('地址修改成功');
      setAddrModalVisible(false);
    } catch {
      message.error('地址修改失败');
    } finally {
      setAddrSaving(false);
    }
  };
  const qInfo = editingAnswer?.qInfo as QuestionInfo | undefined;
  const options = qInfo?.options || [];
  const hasOtherOption = options.some((o) => o.inputs === 2);
  const otherOption = options.find((o) => o.inputs === 2);
  const isOtherSelected = otherOption
    ? (editingAnswer?.qType === 'single_choice' ? editSelected === otherOption.content : editMultiArr.includes(otherOption.content || ''))
    : false;
  const maxSelect = qInfo?.valueNumber || 999;
  const isMultiMaxed = (editingAnswer?.qType === 'multiple_choice') && editMultiArr.length >= maxSelect;

  const renderEditControl = () => {
    const qType = editingAnswer?.qType;

    // 单选题
    if (qType === 'single_choice') {
      return (
        <Radio.Group value={editSelected} onChange={(e) => setEditSelected(e.target.value)} style={{ display: 'flex', flexDirection: 'column', gap: 8 }}>
          {options.map((opt, i) => (
            <Radio key={i} value={opt.content}>
              {opt.content}{opt.inputs === 2 ? '...' : ''}
              {opt.inputs === 2 && editSelected === opt.content && (
                <Input
                  value={editOtherText}
                  onChange={(e) => setEditOtherText(e.target.value)}
                  placeholder="请输入...(必填)"
                  status={isOtherSelected && !editOtherText.trim() ? 'error' : undefined}
                  style={{ width: 160, marginLeft: 12 }}
                  onClick={(e) => e.stopPropagation()}
                />
              )}
            </Radio>
          ))}
        </Radio.Group>
      );
    }

    // 多选题
    if (qType === 'multiple_choice') {
      return (
        <Checkbox.Group value={editMultiArr} onChange={(vals: any) => setEditMultiArr(vals)} style={{ display: 'flex', flexDirection: 'column', gap: 8 }}>
          {options.map((opt, i) => {
            const disabled = isMultiMaxed && !editMultiArr.includes(opt.content || '');
            return (
              <Checkbox key={i} value={opt.content} disabled={disabled}>
                {opt.content}{opt.inputs === 2 ? '...' : ''}
                {opt.inputs === 2 && editMultiArr.includes(opt.content || '') && (
                  <Input
                    value={editOtherText}
                    onChange={(e) => setEditOtherText(e.target.value)}
                    placeholder="请输入...(必填)"
                    status={isOtherSelected && !editOtherText.trim() ? 'error' : undefined}
                    style={{ width: 160, marginLeft: 12 }}
                    onClick={(e) => e.stopPropagation()}
                  />
                )}
              </Checkbox>
            );
          })}
        </Checkbox.Group>
      );
    }

    // 下拉题
    if (qType === 'dropdown') {
      return (
        <Select value={editSelected || undefined} onChange={(v) => setEditSelected(v)} placeholder="请选择..." style={{ width: '100%' }}>
          {options.map((opt, i) => (
            <Select.Option key={i} value={opt.content}>{opt.content}</Select.Option>
          ))}
        </Select>
      );
    }

    // 数字类
    if (qType === 'number' || qType === 'slider' || qType === 'rate' || qType === 'nps') {
      return <InputNumber value={editNumValue} onChange={(v) => setEditNumValue(v)} style={{ width: '100%' }} />;
    }

    // 日期
    if (qType === 'date') {
      return (
        <DatePicker
          value={editSelected ? dayjs(editSelected) : null}
          onChange={(_, ds) => setEditSelected(ds)}
          style={{ width: '100%' }}
        />
      );
    }

    // 日期时间
    if (qType === 'datetime') {
      return (
        <DatePicker
          showTime
          value={editSelected ? dayjs(editSelected) : null}
          onChange={(_, ds) => setEditSelected(ds)}
          style={{ width: '100%' }}
        />
      );
    }

    // 文本类
    if (qType === 'text' || qType === 'input') {
      return <TextArea value={editOtherText} onChange={(e) => setEditOtherText(e.target.value)} placeholder="输入文本" rows={4} />;
    }

    // 默认
    return <Input value={editSelected} onChange={(e) => setEditSelected(e.target.value)} placeholder="输入答案" />;
  };

  const shownAnswer = (item: any) => {
    const qInfo = questionMap[item.surveyQuestionId];
    const qType = qInfo?.type;
    if (qType === 'text' || qType === 'input') {
      return <span>{item.answerText || '无'}</span>;
    }
    return <span>{item.answer?.join('，') || '无'}</span>;
  };

  return (
    <PageContainer>
      <Card bordered={false}>
        <Descriptions title="信息" style={{ marginBottom: 32 }} column={{ xxl: 3, xl: 3, lg: 2, md: 2, sm: 1, xs: 1 }}>
          <Descriptions.Item label="编号">{response?.sn || ''}</Descriptions.Item>
          <Descriptions.Item label="受访人">{response?.respondent || ''}</Descriptions.Item>
          <Descriptions.Item label="受访人联系电话">{response?.respondentPhone || ''}</Descriptions.Item>
          <Descriptions.Item label="调研员">{response?.researcher || ''}</Descriptions.Item>
          <Descriptions.Item label="调研员联系电话">{response?.researcherPhone || ''}</Descriptions.Item>
          <Descriptions.Item label="合照照片">
            {response?.pic?.map((v) => <Image key={v} width={20} src={v} preview={{ src: v }} />) || ''}
          </Descriptions.Item>
          <Descriptions.Item label="ip">{response?.ip || ''}</Descriptions.Item>
          <Descriptions.Item label="地址">
            <span style={{ wordBreak: 'break-all' }}>
              {response?.area}{response?.city}{response?.district}{response?.village}{response?.address}
            </span>
            {isInterviewer && (
              <div style={{ marginTop: 4 }}>
                <Button type="link" icon={<EditOutlined />} onClick={async () => {
                setAddrProvince(null);
                setAddrCity(null);
                setAddrDistrict(null);
                setAddrTown(null);
                setAddrDetail(response?.address || '');
                setAddrModalVisible(true);

                // 尝试回显省
                const { data: provinces } = await queryProvince();
                const pOpts = (provinces || []).map((item: any) => ({ label: item.title, value: item.value }));
                setProvinceOpts(pOpts);
                const matchedP = pOpts.find((o: any) => o.label === response?.area) || null;
                if (matchedP) {
                  setAddrProvince(matchedP);
                  // 回显市
                  const { data: cities } = await queryCity(matchedP.value);
                  const cOpts = (cities || []).map((item: any) => ({ label: item.title, value: item.value }));
                  setCityOpts(cOpts);
                  const matchedC = cOpts.find((o: any) => o.label === response?.city) || null;
                  if (matchedC) {
                    setAddrCity(matchedC);
                    // 回显县
                    const { data: districts } = await queryCity(matchedC.value);
                    const dOpts = (districts || []).map((item: any) => ({ label: item.title, value: item.value }));
                    setDistrictOpts(dOpts);
                    const matchedD = dOpts.find((o: any) => o.label === response?.district) || null;
                    if (matchedD) {
                      setAddrDistrict(matchedD);
                      // 回显镇
                      const { data: towns } = await queryCity(matchedD.value);
                      const tOpts = (towns || []).map((item: any) => ({ label: item.title, value: item.value }));
                      setTownOpts(tOpts);
                      const matchedT = tOpts.find((o: any) => o.label === response?.village) || null;
                      if (matchedT) setAddrTown(matchedT);
                    }
                  }
                }
              }}>
                编辑地址
              </Button>
              </div>
            )}
          </Descriptions.Item>
          <Descriptions.Item label="地图地址">{response?.latitude}-{response?.longitude}</Descriptions.Item>
          <Descriptions.Item label="回答数量">{response?.answerCount || ''}</Descriptions.Item>
        </Descriptions>
        <Divider style={{ marginBottom: 32 }} />
        <div className={styles.title}>
          详情 {isInterviewer && <span style={{ color: '#1890ff', fontSize: 14 }}>（编辑模式）</span>}
        </div>

        <Card style={{ marginBottom: 24 }}>
          <Steps
            progressDot
            current={responseAnswers.length + 1}
            direction="vertical"
            items={responseAnswers.map((item) => {
              const qInfo = questionMap[item.surveyQuestionId || 0];
              const hasOther = qInfo?.options?.some((o: any) => o.inputs === 2);
              const maxHint = qInfo?.type === 'multiple_choice' && qInfo?.valueNumber
                ? `（最多选${qInfo.valueNumber}项）`
                : '';
              return {
                title: <span>{item.content} {maxHint}</span>,
                description: (
                  <Card>
                    <ul>
                      <li>回答内容: {shownAnswer(item)}</li>
                      {hasOther && <li>其他/补充: {item.answerText || '无'}</li>}
                      {!hasOther && item.answerText && <li>补充说明: {item.answerText}</li>}
                      <li>回答时间: {item.createdAt}</li>
                      {isInterviewer && (
                        <li>
                          <Button type="link" icon={<EditOutlined />} onClick={() => handleEdit(item)}>
                            编辑
                          </Button>
                        </li>
                      )}
                    </ul>
                  </Card>
                ),
              };
            })}
          />
        </Card>
      </Card>

      <Modal
        title={`修改答案 — ${editingAnswer?.content || ''}`}
        open={editModalVisible}
        onOk={handleSave}
        onCancel={() => setEditModalVisible(false)}
        okText="保存"
        cancelText="取消"
        width={520}
      >
        <div style={{ marginBottom: 8, color: '#888' }}>类型: {qTypeLabel(editingAnswer?.qType)}</div>
        <div style={{ marginBottom: 16 }}>
          <div style={{ marginBottom: 8 }}>回答内容：</div>
          {renderEditControl()}
        </div>
        {hasOtherOption && (
          <div style={{ marginBottom: 16 }}>
            <div style={{ marginBottom: 8 }}>其他/补充说明：</div>
            <TextArea
              value={isOtherSelected ? editOtherText : ''}
              onChange={(e) => isOtherSelected && setEditOtherText(e.target.value)}
              placeholder={isOtherSelected ? '请输入补充说明' : '请先选择"其他"选项'}
              rows={3}
              disabled={!isOtherSelected}
            />
          </div>
        )}
      </Modal>

      <Modal
        title="编辑地址"
        open={addrModalVisible}
        onOk={handleAddrSave}
        onCancel={() => setAddrModalVisible(false)}
        okText="保存"
        cancelText="取消"
        confirmLoading={addrSaving}
        width={480}
      >
        <div style={{ marginBottom: 12 }}>
          <div style={{ marginBottom: 4 }}>省（自治区、直辖市）</div>
          <Select
            showSearch
            placeholder="请选择省份"
            style={{ width: '100%' }}
            labelInValue
            value={addrProvince}
            onChange={(v: any) => { setAddrProvince(v); setAddrCity(null); setAddrDistrict(null); setAddrTown(null); setCityOpts([]); setDistrictOpts([]); setTownOpts([]); }}
            options={provinceOpts}
          />
        </div>
        <div style={{ marginBottom: 12 }}>
          <div style={{ marginBottom: 4 }}>市（自治州、地区、盟）</div>
          <Select
            showSearch
            placeholder="请选择城市"
            style={{ width: '100%' }}
            labelInValue
            value={addrCity}
            disabled={!addrProvince}
            options={cityOpts}
            onChange={(v: any) => { setAddrCity(v); setAddrDistrict(null); setAddrTown(null); setDistrictOpts([]); setTownOpts([]); }}
            onDropdownVisibleChange={(open) => {
              if (open && addrProvince?.value && cityOpts.length === 0) {
                queryCity(addrProvince.value).then(({ data }) =>
                  setCityOpts((data || []).map((item: any) => ({ label: item.title, value: item.value }))),
                );
              }
            }}
          />
        </div>
        <div style={{ marginBottom: 12 }}>
          <div style={{ marginBottom: 4 }}>县（自治县、县级市、区）</div>
          <Select
            showSearch
            placeholder="请选择区县"
            style={{ width: '100%' }}
            labelInValue
            value={addrDistrict}
            disabled={!addrCity}
            options={districtOpts}
            onChange={(v: any) => { setAddrDistrict(v); setAddrTown(null); setTownOpts([]); }}
            onDropdownVisibleChange={(open) => {
              if (open && addrCity?.value && districtOpts.length === 0) {
                queryCity(addrCity.value).then(({ data }) =>
                  setDistrictOpts((data || []).map((item: any) => ({ label: item.title, value: item.value }))),
                );
              }
            }}
          />
        </div>
        <div style={{ marginBottom: 12 }}>
          <div style={{ marginBottom: 4 }}>镇（乡、街道）</div>
          <Select
            showSearch
            placeholder="请选择乡镇"
            style={{ width: '100%' }}
            labelInValue
            value={addrTown}
            disabled={!addrDistrict}
            options={townOpts}
            onChange={(v: any) => setAddrTown(v)}
            onDropdownVisibleChange={(open) => {
              if (open && addrDistrict?.value && townOpts.length === 0) {
                queryCity(addrDistrict.value).then(({ data }) =>
                  setTownOpts((data || []).map((item: any) => ({ label: item.title, value: item.value }))),
                );
              }
            }}
          />
        </div>
        <div style={{ marginBottom: 12 }}>
          <div style={{ marginBottom: 4 }}>村/详细地址</div>
          <Input
            value={addrDetail}
            onChange={(e) => setAddrDetail(e.target.value)}
            placeholder="请输入详细地址"
          />
        </div>
      </Modal>
    </PageContainer>
  );
};

export default Basic;