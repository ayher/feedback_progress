// import axios from 'axios';
import configs from '../../config/config';
// import { getTokenHeader } from '../../public/tool/setHead';
import actions from './index';

// const baseUrl = configs.baseUrl;

/**
 * 获取用户形式
 */
 export function getMemberInfo(query = '') {
    return async (dispatch) => {
        console.log("====getMemberInfo")
        try {
            // let value = (JSON.stringify(query));
            // let headers = Object.assign(getTokenHeader({}), {'Content-Type': 'application/json'});
            // const data = (await axios.post(`${baseUrl}/get_member`,
            //     value,
            //     {headers}
            // )).data;
            var data= {
                "name": "默默",
                "wechat": "ayher",
                "avatar": "https://img2.baidu.com/it/u=2212383468,857153027&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500",
                "introduce": "加油！",
                "add_time": "2022-11-26T13:42:46+08:00"
            }

            dispatch({
                type: actions.GET_MEMBER_INFO_SUCCESS,
                data: data
            });
        } catch (error) {
            dispatch({
                type: actions.GET_MEMBER_INFO_FAILURE,
                error: new Error('用户基本信息获取失败, 请稍后再试！')
            });
        }
    };
}