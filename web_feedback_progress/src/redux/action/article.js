import axios from 'axios';
import configs from '../../config/config';
import { getTokenHeader } from '../../public/tool/setHead';
import actions from './index';

const baseUrl = configs.baseUrl;

/**
 * 获取文章形式
 */
export function getArticleInfo(query = '') {
    return async (dispatch) => {
        try {
            let value = (JSON.stringify(query));
            let headers = Object.assign(getTokenHeader({}), {'Content-Type': 'application/json'});
            const data = (await axios.post(`${baseUrl}/get_article`,
                value,
                {headers}
            )).data;

            dispatch({
                type: actions.GET_ARTICLE_INFO_SUCCESS,
                data: data
            });
        } catch (error) {
            dispatch({
                type: actions.GET_ARTICLE_INFO_FAILURE,
                error: new Error('文章基本信息获取失败, 请稍后再试！')
            });
        }
    };
}

/**
 * 获取文章排名
 */
export function getRankArticleInfo(query = '') {
    return async (dispatch) => {
        try {
            let value = (JSON.stringify(query));
            let headers = Object.assign(getTokenHeader({}), {'Content-Type': 'application/json'});
            const data = (await axios.post(`${baseUrl}/get_rank_article`,
                value,
                {headers}
            )).data;

            dispatch({
                type: actions.GET_ARTICLE_RANK_SUCCESS,
                data: data
            });
        } catch (error) {
            dispatch({
                type: actions.GET_ARTICLE_RANK_FAILURE,
                error: new Error('文章排名信息获取失败, 请稍后再试！')
            });
        }
    };
}