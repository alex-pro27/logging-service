#!/bin/bash
set -e

clickhouse client -n <<-EOSQL
    CREATE TABLE IF NOT EXISTS migrations (
			num Int32,
			status Int16,
			event_date Date
		) ENGINE=MergeTree(event_date, (num), 8192);
EOSQL