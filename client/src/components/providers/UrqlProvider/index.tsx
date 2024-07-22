import React from 'react';
import { Provider } from 'urql';

import { client } from '@/lib/urql';

type Props = {
  children: React.ReactNode;
};

export const UrqlProvider: React.FC<Props> = React.memo(({ children }) => {
  return <Provider value={client}>{children}</Provider>;
});
