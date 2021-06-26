import csv

''' 
Read all rows from original csv file and generate a new csv file with only cities in GB
We are interest only in columns: name, latitude, longitude, country code.
The rest can be discarded
'''

with open('cities15000.csv') as csv_file:
    csv_reader = csv.reader(csv_file, delimiter=',')
    with open('cities_cleaned.csv', 'w') as f:
        csv_writer = csv.writer(f)
        csv_writer.writerow(["Name", "Latitude", "Longitude", "CountryCode"])
        line_count = 0
        for row in csv_reader:
            if row[8] == 'GB':
                line_count += 1
                csv_writer.writerow([line_count, row[1], row[4], row[5], row[8]])
    print(f'Processed {line_count} cities in GB.')