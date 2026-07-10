import React, { useState, useRef, useEffect } from 'react';
import { Input, Form } from 'antd';
import { ProFormSelect } from '@ant-design/pro-components';
import { queryProvince, queryCity } from '@/services/ant-design-pro/api';
import QJumpRules from '@/pages/survey/respondent/components/QJumpRules';
import type { QuestionComponentProps } from '@/pages/survey/respondent/types';

const QAddress = (props: QuestionComponentProps) => {
  const { surveyId, question, generateRandom, addRespondent, setCurrentNum, setCurrent } = props;
  const addrRef = useRef<Record<string, string>>({});
  const form = Form.useFormInstance();

  const [provinceVal, setProvinceVal] = useState<{ label: string; value: string } | null>(null);
  const [cityVal, setCityVal] = useState<{ label: string; value: string } | null>(null);
  const [districtVal, setDistrictVal] = useState<{ label: string; value: string } | null>(null);
  const [hasAddressData, setHasAddressData] = useState(true);

  // Track whether API returns empty data for district/town → fall back to input
  const [districtNoData, setDistrictNoData] = useState(false);
  const [townNoData, setTownNoData] = useState(false);

  const districtField = ['address', question.id, 'district'];
  const villageField = ['address', question.id, 'village'];

  // Clear form value when switching from Select to Input to avoid [object Object]
  useEffect(() => {
    if (districtNoData) {
      form.setFieldValue(districtField, '');
      form.setFieldValue(villageField, '');
    }
  }, [districtNoData]);

  useEffect(() => {
    if (townNoData) {
      form.setFieldValue(villageField, '');
    }
  }, [townNoData]);

  if (!question) return null;

  const isRequired = question.required === 1 && hasAddressData;

  const savePart = (type: string, val: string) => {
    addRespondent({
      surveyId,
      type,
      questionId: question.id,
      value: [val],
      sn: generateRandom,
    });
    addrRef.current = { ...addrRef.current, [type]: val };
  };

  const resetDistrict = () => {
    setDistrictVal(null);
    setDistrictNoData(false);
    setTownNoData(false);
  };

  const resetTown = () => {
    setTownNoData(false);
  };

  const loadDistrict = async (): Promise<{ label: string; value: string }[]> => {
    if (!cityVal?.value) return [];
    try {
      const { data } = await queryCity(cityVal.value);
      const opts = (data || []).map((item: any) => ({ label: item.title, value: item.value }));
      if (opts.length === 0) setDistrictNoData(true);
      return opts;
    } catch {
      setDistrictNoData(true);
      return [];
    }
  };

  const loadTown = async (): Promise<{ label: string; value: string }[]> => {
    if (!districtVal?.value) return [];
    try {
      const { data } = await queryCity(districtVal.value);
      const opts = (data || []).map((item: any) => ({ label: item.title, value: item.value }));
      if (opts.length === 0) setTownNoData(true);
      return opts;
    } catch {
      setTownNoData(true);
      return [];
    }
  };

  return (
    <>
      <h3>{question.serial ? question.serial + '-' : ''}{question.content}</h3>

      <ProFormSelect
        label="省（自治区、直辖市）"
        width="md"
        name={['address', question.id, 'province']}
        rules={isRequired ? [{ required: true, message: '请选择省份' }] : []}
        fieldProps={{ labelInValue: true }}
        request={async () =>
          queryProvince().then(({ data }) => {
            if (!data?.length) setHasAddressData(false);
            return (data || []).map((item: any) => ({ label: item.title, value: item.value }));
          })
        }
        onChange={(e: any) => {
          setProvinceVal(e);
          setCityVal(null);
          resetDistrict();
          savePart('area', e?.label || '');
        }}
      />

      <ProFormSelect
        label="市（自治州、地区、盟）"
        width="md"
        name={['address', question.id, 'city']}
        rules={isRequired ? [{ required: true, message: '请选择城市' }] : []}
        fieldProps={{ labelInValue: true }}
        disabled={!provinceVal}
        dependencies={[['address', question.id, 'province']]}
        request={async () => {
          if (!provinceVal?.value) return [];
          return queryCity(provinceVal.value).then(({ data }) =>
            (data || []).map((item: any) => ({ label: item.title, value: item.value })),
          );
        }}
        onChange={(e: any) => {
          setCityVal(e);
          resetDistrict();
          savePart('city', e?.label || '');
        }}
      />

      {districtNoData ? (
        <Form.Item
          label="县（自治县、县级市、区）"
          name={['address', question.id, 'district']}
          rules={isRequired ? [{ required: true, message: '请输入区县' }] : []}
        >
          <Input
            placeholder="请输入区县名称"
            onChange={(e) => savePart('district', e.target.value)}
          />
        </Form.Item>
      ) : (
        <ProFormSelect
          label="县（自治县、县级市、区）"
          width="md"
          name={['address', question.id, 'district']}
          rules={isRequired ? [{ required: true, message: '请选择区县' }] : []}
          fieldProps={{ labelInValue: true }}
          disabled={!cityVal}
          dependencies={[['address', question.id, 'city']]}
          request={loadDistrict}
          onChange={(e: any) => {
            setDistrictVal(e);
            resetTown();
            savePart('district', e?.label || '');
          }}
        />
      )}

      {townNoData || districtNoData ? (
        <Form.Item
          label="镇（乡、街道）"
          name={['address', question.id, 'village']}
          rules={isRequired ? [{ required: true, message: '请输入乡镇' }] : []}
        >
          <Input
            placeholder="请输入乡镇名称"
            onChange={(e) => savePart('village', e.target.value)}
          />
        </Form.Item>
      ) : (
        <ProFormSelect
          label="镇（乡、街道）"
          width="md"
          name={['address', question.id, 'village']}
          rules={isRequired ? [{ required: true, message: '请选择乡镇' }] : []}
          fieldProps={{ labelInValue: true }}
          disabled={!districtVal}
          dependencies={[['address', question.id, 'district']]}
          request={loadTown}
          onChange={(e: any) => savePart('village', e?.label || '')}
        />
      )}

      <Form.Item label="村">
        <Input
          placeholder="请输入村名"
          onChange={(e) => savePart('address', e.target.value)}
        />
      </Form.Item>

      <QJumpRules
        surveyId={surveyId}
        question={question}
        generateRandom={generateRandom}
        addRespondent={addRespondent}
        setCurrentNum={setCurrentNum}
        setCurrent={setCurrent}
        value={addrRef.current}
      />
    </>
  );
};

export default QAddress;