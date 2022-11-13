FROM node:16

WORKDIR /project/web_feedback_progress
CMD ["npm","start"] 
# CMD ["ping","www.baidu.com"] # 需要自己去 npm install