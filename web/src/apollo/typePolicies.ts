import { ResolversTypes } from 'Generated/schema';
import { TypePolicy } from '@apollo/client';
import Storage from 'Helpers/storage';
import { ERROR_REPORTING_CONSENT_STORAGE_KEY } from 'Source/constants';

const typePolicies: Partial<Record<keyof ResolversTypes, TypePolicy>> = {
  Query: {
    fields: {
      getComplianceIntegration(existingData, { args, toReference }) {
        return (
          existingData ||
          toReference({ __typename: 'ComplianceIntegration', integrationId: args.id })
        );
      },
      getLogIntegration(existingData, { args, toReference }) {
        return (
          existingData || toReference({ __typename: 'LogIntegration', integrationId: args.id })
        );
      },
      getLogDatabase(existingData, { args, toReference }) {
        return existingData || toReference({ __typename: 'LogDatabase', name: args.name });
      },
      getLogDatabaseTable(existingData, { args, toReference }) {
        return (
          existingData ||
          toReference({
            __typename: 'LogDatabaseTable',
            name: args.input.name,
            databaseName: args.input.databaseName,
          })
        );
      },
    },
  },
  Destination: {
    keyFields: ['outputId'],
  },
  AlertDetails: {
    keyFields: ['alertId'],
  },
  AlertSummary: {
    keyFields: ['alertId'],
  },
  ComplianceIntegration: {
    keyFields: ['integrationId'],
  },
  LogIntegration: {
    keyFields: ['integrationId'],
  },
  LogDatabase: {
    keyFields: ['name'],
  },
  LogDatabaseTable: {
    keyFields: ['name'],
  },
  GeneralSettings: {
    keyFields: ['email'],
    fields: {
      errorReportingConsent: {
        merge(oldValue, newValue) {
          Storage.write(ERROR_REPORTING_CONSENT_STORAGE_KEY, newValue);
          return newValue;
        },
      },
    },
  },
};

export default typePolicies;
