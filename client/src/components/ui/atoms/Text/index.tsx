import React from 'react';

import type { Color } from '@/theme/color';
import type { FontSize } from '@/theme/fontSize';
import type { FontWeight } from '@/theme/fontWeight';

import { styles } from './style.css';

type StyleProps = {
  color?: Color;
  size?: FontSize;
  weight?: FontWeight;
};

type Props = {
  children: string;
} & StyleProps;

export const Text: React.FC<Props> = React.memo(
  ({ children, color = 'black', size = 'm', weight = 'normal' }) => {
    return <p className={styles({ color, size, weight })}>{children}</p>;
  },
);
