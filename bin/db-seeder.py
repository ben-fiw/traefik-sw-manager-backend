import random
import uuid
from datetime import datetime, timedelta

# Function to generate random timestamp within a given range
def random_timestamp(start_date, end_date):
    delta = end_date - start_date
    random_days = random.randint(0, delta.days)
    random_seconds = random.randint(0, delta.seconds)
    return start_date + timedelta(days=random_days, seconds=random_seconds)

# Function to generate random UUID
def generate_uuid():
    return str(uuid.uuid4())

# Function to generate random IPv4 address
def generate_ip_address():
    return ".".join(str(random.randint(0, 255)) for _ in range(4))

# Function to generate random IP address
def generate_ip_address():
    return ".".join(str(random.randint(0, 255)) for _ in range(4))

# Function to generate random hex string
def generate_hex_string(length):
    return ''.join(random.choice('0123456789abcdef') for _ in range(length))

# Function to insert test data into available_versions table
def insert_versions(cursor, versions):
    for version in versions:
        display_name = f"Shopware {version}"
        created_at = random_timestamp(datetime.now() - timedelta(days=365), datetime.now())
        updated_at = random_timestamp(created_at, datetime.now())
        cursor.execute("INSERT INTO available_versions (id, version, display_name, created_at, updated_at) VALUES (%s, %s, %s, %s, %s)",
                (generate_uuid(), version, display_name, created_at, updated_at))

# Define data parameters
versions = [
    "6.6.1.0",
    "6.6.0.2",
    "6.6.0.1",
    "6.6.0.0",
    "6.5.8.9",
    "6.5.8.8",
    "6.5.8.7",
    "6.5.8.6",
    "6.5.8.5",
    "6.5.8.4",
    "6.5.8.3",
    "6.5.8.2",
    "6.5.8.1",
    "6.5.8.0",
    "6.5.7.4",
    "6.5.7.3",
    "6.5.7.2",
    "6.5.7.1",
    "6.5.7.0",
    "6.5.6.1",
    "6.5.6.0",
    "6.5.5.2",
    "6.5.5.1",
    "6.5.5.0",
    "6.5.4.1",
    "6.5.4.0",
    "6.5.3.3",
    "6.5.3.2",
    "6.5.3.1",
    "6.5.3.0",
    "6.5.2.1",
    "6.5.2.0",
    "6.5.1.1",
    "6.5.1.0",
    "6.5.0.0"
]

# Connect to your database using appropriate library (e.g., psycopg2 for PostgreSQL, mysql.connector for MySQL)
# Replace the placeholders with actual connection details
# conn = psycopg2.connect(host="your_host", user="your_user", password="your_password", database="your_database")
# cursor = conn.cursor()

# Example MySQL connection using mysql.connector
import mysql.connector

conn = mysql.connector.connect(host="localhost", user="root", password="root", database="demo_shop_manager")
cursor = conn.cursor()

# Insert test data into service_connections table
insert_versions(cursor, versions)

# Commit changes and close connection
conn.commit()
cursor.close()
conn.close()
