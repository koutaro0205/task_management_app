import type { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
  overwrite: true,
  schema: './graphql/schema/*.graphql',
  documents: 'src/**/*.ts',
  // emitLegacyCommonJSImports: false,
  generates: {
    './src/gql/': {
      preset: 'client',
      plugins: [],
      presetConfig: {
        fragmentMasking: false,
      },
    },
    // './graphql.schema.json': {
    //   plugins: ['introspection'],
    // },
  },
  watch: true,
};

export default config;
