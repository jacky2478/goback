#pragma once

// ͨѶ���������ֶ� Type, Data�������ڲ�������ҵ��ָ��
// 1��DUIToBack��Ҫָ������״̬������Init,Run,Over,�ٸ�����������
// ʾ��:{"Type":"Init", Data:""}, {"Type":"Run", Data:""}, {"Type":"Over", Data:""}
// 2��BackToDUI��ҵ�������������֣���������UI���ݣ�������־���ݣ�����ϵͳ��Ϣ���ٸ�����������
// ʾ��:{"Type":"Edit", "Data":""}, {"Type":"Log", Data:""}, {"Type":"Sys", Data:""}

namespace Service
{
	namespace Net
	{
		typedef void callback_func_type(CString data);
		class Comm
		{
		public:
			static bool BackToDUI(CString strData);
			static bool DUIToBack(CString strData);
			static void Init(callback_func_type *pFunc);

			static callback_func_type *s_pFunc;
		};
	}
}

