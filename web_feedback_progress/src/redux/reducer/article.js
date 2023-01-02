import actions from '../action/index';

const {
    GET_ARTICLE_INFO_SUCCESS,
    GET_ARTICLE_RANK_SUCCESS

} = actions;
export default (state = {}, action) => {
    switch (action.type) {
        case GET_ARTICLE_INFO_SUCCESS:
            return {
                ...state,
                ArticleInfo: action.data,
            }
        case GET_ARTICLE_RANK_SUCCESS:
            return {
                ...state,
                ArticleRand: action.data,
            }
        default:
            return state;
    }
};