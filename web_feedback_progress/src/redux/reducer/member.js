import actions from '../action/index';

const {
    GET_MEMBER_INFO_SUCCESS,
} = actions;
export default (state = {}, action) => {
    switch (action.type) {
        case GET_MEMBER_INFO_SUCCESS:
            return {
                ...state,
                MemberInfo: action.data,
            }
        default:
            return state;
    }
};