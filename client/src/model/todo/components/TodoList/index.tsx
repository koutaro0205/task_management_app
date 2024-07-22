import React from 'react';
import { useQuery } from 'urql';

import { GET_TODOS } from './graphql';

export const TodoList: React.FC = React.memo(() => {
  const [result] = useQuery({
    query: GET_TODOS,
  });

  console.log(JSON.stringify(result.data, null, '\t'));
  // eslint-disable-next-line prettier/prettier
  return <p>コンソールを開いて取得されたデータを確認してください</p>;
});
